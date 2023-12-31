package application

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
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
		return "", &response.Error{
			Code:    404,
			Message: "Username atau Password salah",
		}
	}

	data := map[string]interface{}{
		"id":   result.ID,
		"role": role,
	}

	// create jwt token
	jwt, err := utils.GenerateToken(data)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (s *AuthService) VerifyToken(jwt string) (*jwt.MapClaims, error) {
	// find user by email from database
	mapClaims, err := utils.VerifyToken(jwt)

	if err != nil {
		return nil, err
	}

	return &mapClaims, nil
}
