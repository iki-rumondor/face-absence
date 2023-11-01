package application

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
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

func (s *AuthService) VerifyUser(user *domain.User) (string, error) {
	// find user by email from database
	result, err := s.Repo.FindByEmail(user.Email)
	if err != nil {
		return "", errors.New("sorry, the provided email is not registered in our system")
	}

	// verify user password
	if err := utils.ComparePassword(result.Password, user.Password); err != nil {
		return "", errors.New("whoops! password mismatch")
	}

	data := map[string]interface{}{
		"id": result.ID,
		"role": result.RoleID,
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
