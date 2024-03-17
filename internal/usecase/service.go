package usecase

import (
	"github.com/nanmenkaimak/film_library/internal/config"
	"github.com/nanmenkaimak/film_library/internal/repository"
	"go.uber.org/zap"
)

type Service struct {
	repo         repository.Repository
	jwtSecretKey string
	logger       *zap.SugaredLogger
}

func NewService(repo repository.Repository, config config.Auth, logger *zap.SugaredLogger) UseCase {
	return &Service{
		repo:         repo,
		jwtSecretKey: config.JwtSecretKey,
		logger:       logger,
	}
}
