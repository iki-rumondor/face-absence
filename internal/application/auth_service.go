package application

import (
	"errors"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	Repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{
		Repo: repo,
	}
}

func (s *AuthService) VerifyUser(role string, user *domain.User) (string, error) {

	// find user from database
	result, err := s.Repo.FindByUsername(user.Username)
	if err != nil {
		log.Println("Username Salah")
		return "", &response.Error{
			Code:    404,
			Message: "Username atau Password salah",
		}
	}

	// check user role
	if role == "GURU" {
		if err := s.Repo.FindTeacherByUserID(result.ID); err != nil {
			return "", &response.Error{
				Code:    404,
				Message: "Guru dengan username tersebut tidak ditemukan",
			}
		}
	}
	if role == "SANTRI" {
		if err := s.Repo.FindStudentByUserID(result.ID); err != nil {
			return "", &response.Error{
				Code:    404,
				Message: "Santri dengan username tersebut tidak ditemukan",
			}
		}
	}
	if role == "ADMIN" {
		if err := s.Repo.FindAdminByUserID(result.ID); err != nil {
			return "", &response.Error{
				Code:    404,
				Message: "Admin dengan username tersebut tidak ditemukan",
			}
		}
	}

	// verify user password
	if err := utils.ComparePassword(result.Password, user.Password); err != nil {
		log.Println("Password salah")
		return "", &response.Error{
			Code:    404,
			Message: "Username atau Password salah",
		}
	}

	// create jwt token
	jwt, err := utils.GenerateToken(result.ID, role)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (s *AuthService) VerifyToken(jwt string) (*jwt.MapClaims, error) {
	mapClaims, err := utils.VerifyToken(jwt)

	if err != nil {
		return nil, err
	}

	return &mapClaims, nil
}

func (s *AuthService) GetUserByID(id uint) (*domain.User, error) {
	user, err := s.Repo.FindUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.Error{
				Code:    404,
				Message: fmt.Sprintf("User dengan id %d tidak ditemukan", id),
			}
		}
		return nil, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	return user, nil
}
