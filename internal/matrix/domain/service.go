package domain

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/google/uuid"
)

type ServiceInterface interface {
	Matrix(ctx context.Context, id uuid.UUID) (Matrix, error)
	Matrices(ctx context.Context) ([]Matrix, error)
	CreateMatrix(ctx context.Context, matrix *Matrix) error
	UpdateMatrix(ctx context.Context, matrix *Matrix) error
	DeleteMatrix(ctx context.Context, id uuid.UUID) error
	AddSubject(ctx context.Context, matrixID, subjectID uuid.UUID) error
	RemoveSubject(ctx context.Context, matrixID, SubjectID uuid.UUID) error
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

func (s *Service) AddSubject(_ context.Context, matrixID, subjectID uuid.UUID) error {
	if err := s.repo.AddSubject(matrixID, subjectID); err != nil {
		return fmt.Errorf("service can't adds the subject to matrix: %w", err)
	}
	return nil
}

func (s *Service) RemoveSubject(_ context.Context, matrixID, SubjectID uuid.UUID) error {
	if err := s.repo.RemoveSubject(matrixID, SubjectID); err != nil {
		return fmt.Errorf("service can't removes the subject from matrix: %w", err)
	}
	return nil
}
