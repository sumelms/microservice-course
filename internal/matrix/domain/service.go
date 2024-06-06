package domain

import (
	"context"

	"github.com/go-kit/log"
	"github.com/google/uuid"
)

type ServiceInterface interface {
	Matrix(ctx context.Context, matrixUUID uuid.UUID) (Matrix, error)
	MatrixSubject(ctx context.Context, matrixSubject *MatrixSubject) error
	CourseMatrixExists(ctx context.Context, courseUUID uuid.UUID, matrixUUID uuid.UUID) (bool, error)
	Matrices(ctx context.Context, filters *MatrixFilters) ([]Matrix, error)
	MatrixSubjects(ctx context.Context, matrixUUID uuid.UUID) ([]MatrixSubject, error)
	CreateMatrix(ctx context.Context, matrix *Matrix) error
	UpdateMatrix(ctx context.Context, matrix *Matrix) error
	DeleteMatrix(ctx context.Context, matrix *DeletedMatrix) error
	CreateMatrixSubject(ctx context.Context, matrixSubject *MatrixSubject) error
	UpdateMatrixSubject(ctx context.Context, matrixSubject *MatrixSubject) error
	DeleteMatrixSubject(ctx context.Context, matrixSubject *DeletedMatrixSubject) error

	Subject(ctx context.Context, subjectUUID uuid.UUID) (Subject, error)
	Subjects(ctx context.Context) ([]Subject, error)
	CreateSubject(ctx context.Context, subject *Subject) error
	UpdateSubject(ctx context.Context, subject *Subject) error
	DeleteSubject(ctx context.Context, subject *DeletedSubject) error
}

type ServiceConfiguration func(svc *Service) error

type Service struct {
	matrices MatrixRepository
	subjects SubjectRepository
	courses  CourseClient
	logger   log.Logger
}

// NewService creates a new domain Service instance.
func NewService(cfgs ...ServiceConfiguration) (*Service, error) {
	svc := &Service{}
	for _, cfg := range cfgs {
		err := cfg(svc)
		if err != nil {
			return nil, err
		}
	}

	return svc, nil
}

// WithMatrixRepository injects the course repository to the domain Service.
func WithMatrixRepository(cr MatrixRepository) ServiceConfiguration {
	return func(svc *Service) error {
		svc.matrices = cr

		return nil
	}
}

// WithSubjectRepository injects the subscription repository to the domain Service.
func WithSubjectRepository(sr SubjectRepository) ServiceConfiguration {
	return func(svc *Service) error {
		svc.subjects = sr

		return nil
	}
}

// WithLogger injects the logger to the domain Service.
func WithLogger(l log.Logger) ServiceConfiguration {
	return func(svc *Service) error {
		svc.logger = l

		return nil
	}
}

func WithCourseClient(courses CourseClient) ServiceConfiguration {
	return func(svc *Service) error {
		svc.courses = courses

		return nil
	}
}
