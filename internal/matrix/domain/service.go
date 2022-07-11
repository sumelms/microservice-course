package domain

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/google/uuid"
)

type ServiceInterface interface {
	Matrix(context.Context, uuid.UUID) (Matrix, error)
	Matrices(context.Context) ([]Matrix, error)
	CreateMatrix(context.Context, *Matrix) error
	UpdateMatrix(context.Context, *Matrix) error
	DeleteMatrix(context.Context, uuid.UUID) error
}

type Service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

func (s *Service) Matrix(_ context.Context, id uuid.UUID) (Matrix, error) {
	m, err := s.repo.Matrix(id)
	if err != nil {
		return Matrix{}, fmt.Errorf("service can't find matrix: %w", err)
	}
	return m, nil
}

func (s *Service) Matrices(_ context.Context) ([]Matrix, error) {
	mm, err := s.repo.Matrices()
	if err != nil {
		return []Matrix{}, fmt.Errorf("service didn't found any matrix: %w", err)
	}
	return mm, nil
}

func (s *Service) CreateMatrix(_ context.Context, m *Matrix) error {
	if err := s.repo.CreateMatrix(m); err != nil {
		return fmt.Errorf("service can't create matrix: %w", err)
	}
	return nil
}

func (s *Service) UpdateMatrix(_ context.Context, m *Matrix) error {
	if err := s.repo.UpdateMatrix(m); err != nil {
		return fmt.Errorf("service can't update matrix: %w", err)
	}
	return nil
}

func (s *Service) DeleteMatrix(_ context.Context, id uuid.UUID) error {
	if err := s.repo.DeleteMatrix(id); err != nil {
		return fmt.Errorf("service can't delete matrix: %w", err)
	}
	return nil
}
