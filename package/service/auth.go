package service

import (
	"crypto/sha1"
	"fmt"
	"toDo"
	"toDo/package/repository"
)

const salt = "fdsnjsfgndfjngdfjgn331123"

type AuthService struct {
	repo repository.Authorization
}

// NewAuthService this is constructor
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s AuthService) CreateUser(user toDo.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
