package learning_path

import "log"

type ServiceInterface interface {
}

type Service struct {
	repo   RepositoryInterface
	logger log.Logger
}

func NewService(repository RepositoryInterface, logger log.Logger) *Service {
	return &Service{
		repo:   repository,
		logger: logger,
	}
}
