package domain

import "github.com/go-kit/kit/log"

type Service interface{}

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) *service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}
