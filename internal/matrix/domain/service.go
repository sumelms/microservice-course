package domain

import (
	"context"

	"github.com/go-kit/log"
	"github.com/google/uuid"
)

type ServiceInterface interface {
	Matrix(ctx context.Context, matrixUUID uuid.UUID) (Matrix, error)
	CourseMatrix(ctx context.Context, courseUUID uuid.UUID, matrixUUID uuid.UUID) (Matrix, error)
	Matrices(ctx context.Context, filters *MatrixFilters) ([]Matrix, error)
	CreateMatrix(ctx context.Context, matrix *Matrix) error
	UpdateMatrix(ctx context.Context, matrix *Matrix) error
	DeleteMatrix(ctx context.Context, matrixUUID uuid.UUID) error
	AddSubject(ctx context.Context, matrixSubject *MatrixSubject) error
	RemoveSubject(ctx context.Context, matrixID, SubjectID uuid.UUID) error

	Subject(ctx context.Context, id uuid.UUID) (Subject, error)
	Subjects(ctx context.Context) ([]Subject, error)
	CreateSubject(ctx context.Context, subject *Subject) error
	UpdateSubject(ctx context.Context, subject *Subject) error
	DeleteSubject(ctx context.Context, id uuid.UUID) error
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

func WithCourseClient(c CourseClient) ServiceConfiguration {
	return func(svc *Service) error {
		svc.courses = c

		return nil
	}
}
