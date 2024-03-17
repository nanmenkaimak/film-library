package usecase

import (
	"context"
	"fmt"
	"github.com/nanmenkaimak/film_library/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

func (a *Service) RegisterUser(ctx context.Context, request entity.User) (*RegisterUserResponse, error) {
	hashPass, err := a.hashPassword(request.Password)
	if err != nil {
		return nil, fmt.Errorf("hashing password err: %v", err)
	}
	request.Password = hashPass

	userID, err := a.repo.CreateUser(request)
	if err != nil {
		return nil, fmt.Errorf("CreateUser err: %w", err)
	}
	resp := &RegisterUserResponse{
		ID: userID,
	}

	return resp, nil
}

func (a *Service) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
