package service

import (
	"crypto/sha1"
	"fmt"

	gotodo "github.com/grancc/go-to-do-app"
	"github.com/grancc/go-to-do-app/pkg/repository"
)

const salt = "nboiwgbOBGRbjjnoiwrjeggntrjfkegnf"

type AutService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AutService {
	return &AutService{repo: repo}
}

func (s *AutService) CreateUser(user gotodo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
