package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) CourseMatrixExists(_ context.Context, courseUUID uuid.UUID, matrixUUID uuid.UUID) (bool, error) {
	exists, err := s.matrices.CourseMatrixExists(courseUUID, matrixUUID)
	if err != nil {
		return false, fmt.Errorf(
			"service can't find matrix with UUID %s from course with UUID %s: %w",
			matrixUUID,
			courseUUID,
			err)
	}
	return exists, nil
}

func (s *Service) Matrix(_ context.Context, matrixUUID uuid.UUID) (Matrix, error) {
	matrix, err := s.matrices.Matrix(matrixUUID)
	if err != nil {
		return Matrix{}, fmt.Errorf("service can't find matrix with UUID %s: %w", matrixUUID, err)
	}
	return matrix, nil
}

func (s *Service) Matrices(_ context.Context, filters *MatrixFilters) ([]Matrix, error) {
	list, err := s.filteredMatrices(filters)
	if err != nil {
		return []Matrix{}, fmt.Errorf("service didn't find any matrices: %w", err)
	}

	return list, nil
}

func (s *Service) filteredMatrices(filters *MatrixFilters) ([]Matrix, error) {
	if filters != nil && filters.CourseUUID != uuid.Nil {
		return s.matrices.CourseMatrices(filters.CourseUUID)
	}
	return s.matrices.Matrices()
}

func (s *Service) CreateMatrix(ctx context.Context, matrix *Matrix) error {
	err := s.courses.CourseExists(ctx, matrix.CourseUUID)
	if err != nil {
		return fmt.Errorf("error checking if course with UUID %s exists: %w", matrix.CourseUUID, err)
	}
	if err := s.matrices.CreateMatrix(matrix); err != nil {
		return fmt.Errorf("service can't create matrix: %w", err)
	}
	return nil
}

func (s *Service) UpdateMatrix(_ context.Context, matrix *Matrix) error {
	err := s.matrices.UpdateMatrix(matrix)
	if err != nil {
		return fmt.Errorf("service can't update matrix: %w", err)
	}
	return nil
}

func (s *Service) DeleteMatrix(_ context.Context, matrix *DeletedMatrix) error {
	if err := s.matrices.DeleteMatrix(matrix); err != nil {
		return fmt.Errorf("service can't delete matrix with UUID %s: %w", matrix.UUID, err)
	}
	return nil
}

func (s *Service) CreateMatrixSubject(_ context.Context, matrixSubject *MatrixSubject) error {
	if err := s.matrices.CreateMatrixSubject(matrixSubject); err != nil {
		return fmt.Errorf("service can't adds the subject to matrix: %w", err)
	}
	return nil
}

func (s *Service) MatrixSubjects(_ context.Context, matrixUUID uuid.UUID) ([]MatrixSubject, error) {
	matrixSubjects, err := s.matrices.MatrixSubjects(matrixUUID)
	if err != nil {
		return []MatrixSubject{}, fmt.Errorf("service can't find matrix with UUID %s: %w", matrixUUID, err)
	}
	return matrixSubjects, nil
}

func (s *Service) MatrixSubject(_ context.Context, matrixSubject *MatrixSubject) error {
	err := s.matrices.MatrixSubject(matrixSubject)
	if err != nil {
		return fmt.Errorf("service can't find matrix subject: %w", err)
	}
	return nil
}

func (s *Service) UpdateMatrixSubject(_ context.Context, matrixSubject *MatrixSubject) error {
	err := s.matrices.UpdateMatrixSubject(matrixSubject)
	if err != nil {
		return fmt.Errorf("service can't update matrix subject: %w", err)
	}
	return nil
}

func (s *Service) DeleteMatrixSubject(_ context.Context, matrixSubject *DeletedMatrixSubject) error {
	if err := s.matrices.DeleteMatrixSubject(matrixSubject); err != nil {
		return fmt.Errorf(
			"service can't remove subject with UUID %s from matrix with UUID %s: %w",
			matrixSubject.SubjectUUID,
			matrixSubject.MatrixUUID,
			err)
	}
	return nil
}
