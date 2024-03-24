package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) CourseMatrix(_ context.Context, courseUUID uuid.UUID, matrixUUID uuid.UUID) (Matrix, error) {
	m, err := s.matrices.CourseMatrix(courseUUID, matrixUUID)
	if err != nil {
		return Matrix{}, fmt.Errorf("service can't find matrix: %w", err)
	}
	return m, nil
}

func (s *Service) Matrix(_ context.Context, matrixUUID uuid.UUID) (Matrix, error) {
	matrix, err := s.matrices.Matrix(matrixUUID)
	if err != nil {
		return Matrix{}, fmt.Errorf("service can't find matrix: %w", err)
	}
	return matrix, nil
}

func (s *Service) Matrices(_ context.Context, filters *MatrixFilters) ([]Matrix, error) {
	list, err := func() ([]Matrix, error) {
		if filters != nil {
			if filters.CourseUUID != uuid.Nil {
				return s.matrices.CourseMatrices(filters.CourseUUID)
			}
		}

		return s.matrices.Matrices()
	}()
	if err != nil {
		return []Matrix{}, fmt.Errorf("service didn't found any matrix: %w", err)
	}

	return list, nil
}

func (s *Service) CreateMatrix(ctx context.Context, matrix *Matrix) error {
	err := s.courses.CourseExists(ctx, *matrix.CourseUUID)
	if err != nil {
		return fmt.Errorf("error checking if course %s exists: %w", matrix.CourseUUID, err)
	}
	if err := s.matrices.CreateMatrix(matrix); err != nil {
		return fmt.Errorf("service can't create matrix: %w", err)
	}
	return nil
}

func (s *Service) UpdateMatrix(_ context.Context, matrix *Matrix) error {
	if err := s.matrices.UpdateMatrix(matrix); err != nil {
		return fmt.Errorf("service can't update matrix: %w", err)
	}
	return nil
}

func (s *Service) DeleteMatrix(_ context.Context, MatrixUUID uuid.UUID) error {
	if err := s.matrices.DeleteMatrix(MatrixUUID); err != nil {
		return fmt.Errorf("service can't delete matrix: %w", err)
	}
	return nil
}

func (s *Service) AddSubject(_ context.Context, matrixSubject *MatrixSubject) error {
	if err := s.matrices.AddSubject(matrixSubject); err != nil {
		return fmt.Errorf("service can't adds the subject to matrix: %w", err)
	}
	return nil
}

func (s *Service) RemoveSubject(_ context.Context, matrixUUID, subjectUUID uuid.UUID) error {
	if err := s.matrices.RemoveSubject(matrixUUID, subjectUUID); err != nil {
		return fmt.Errorf("service can't removes the subject from matrix: %w", err)
	}
	return nil
}
