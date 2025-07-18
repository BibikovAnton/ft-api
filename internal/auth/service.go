package auth

import (
	"errors"

	"github.com/BibikovAnton/finance-tracker-api/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	existedUser, _ := service.userRepository.FindByEmail(email)
	if existedUser != nil {
		return "", errors.New("user exsists")
	}

	hadhedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}

	user := &user.User{
		Email:    email,
		Password: string(hadhedPassword),
		Name:     name,
	}

	service.userRepository.Create(user)
	return user.Email, nil
}

func (service *AuthService) Login(email, password string) (string, error) {
	existedUser, _ := service.userRepository.FindByEmail(email)
	if existedUser == nil {
		return "", errors.New("wrong email or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))

	if err != nil {
		return "", errors.New("wrong email or password")
	}
	return existedUser.Email, nil
}
