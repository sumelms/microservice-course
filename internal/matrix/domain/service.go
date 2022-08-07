package domain

import (
	"context"

	"github.com/go-kit/log"
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

	Subject(context.Context, uuid.UUID) (Subject, error)
	Subjects(context.Context) ([]Subject, error)
	CreateSubject(context.Context, *Subject) error
	UpdateSubject(context.Context, *Subject) error
	DeleteSubject(context.Context, uuid.UUID) error
}

type serviceConfiguration func(svc *Service) error

type Service struct {
	matrices MatrixRepository
	subjects SubjectRepository
	logger   log.Logger
}

// NewService creates a new domain Service instance
func NewService(cfgs ...serviceConfiguration) (*Service, error) {
	svc := &Service{}
	for _, cfg := range cfgs {
		err := cfg(svc)
		if err != nil {
			return nil, err
		}
	}
	return svc, nil
}

// WithMatrixRepository injects the course repository to the domain Service
func WithMatrixRepository(cr MatrixRepository) serviceConfiguration {
	return func(svc *Service) error {
		svc.matrices = cr
		return nil
	}
}

// WithSubjectRepository injects the subscription repository to the domain Service
func WithSubjectRepository(sr SubjectRepository) serviceConfiguration {
	return func(svc *Service) error {
		svc.subjects = sr
		return nil
	}
}

// WithLogger injects the logger to the domain Service
func WithLogger(l log.Logger) serviceConfiguration {
	return func(svc *Service) error {
		svc.logger = l
		return nil
	}
}
