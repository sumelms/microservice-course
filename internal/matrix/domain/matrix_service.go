package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) Matrix(_ context.Context, id uuid.UUID) (Matrix, error) {
	m, err := s.matrices.Matrix(id)
	if err != nil {
		return Matrix{}, fmt.Errorf("service can't find matrix: %w", err)
	}
	return m, nil
}

func (s *Service) Matrices(_ context.Context) ([]Matrix, error) {
	mm, err := s.matrices.Matrices()
	if err != nil {
		return []Matrix{}, fmt.Errorf("service didn't found any matrix: %w", err)
	}
	return mm, nil
}

func (s *Service) CreateMatrix(_ context.Context, m *Matrix) error {
	if err := s.matrices.CreateMatrix(m); err != nil {
		return fmt.Errorf("service can't create matrix: %w", err)
	}
	return nil
}

func (s *Service) UpdateMatrix(_ context.Context, m *Matrix) error {
	if err := s.matrices.UpdateMatrix(m); err != nil {
		return fmt.Errorf("service can't update matrix: %w", err)
	}
	return nil
}

func (s *Service) DeleteMatrix(_ context.Context, id uuid.UUID) error {
	if err := s.matrices.DeleteMatrix(id); err != nil {
		return fmt.Errorf("service can't delete matrix: %w", err)
	}
	return nil
}

func (s *Service) AddSubject(_ context.Context, matrixID, subjectID uuid.UUID) error {
	if err := s.matrices.AddSubject(matrixID, subjectID); err != nil {
		return fmt.Errorf("service can't adds the subject to matrix: %w", err)
	}
	return nil
}

func (s *Service) RemoveSubject(_ context.Context, matrixID, subjectID uuid.UUID) error {
	if err := s.matrices.RemoveSubject(matrixID, subjectID); err != nil {
		return fmt.Errorf("service can't removes the subject from matrix: %w", err)
	}
	return nil
}
