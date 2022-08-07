package domain

import (
	"context"
	"fmt"

	"github.com/go-kit/log"
	"github.com/google/uuid"

	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

type ServiceInterface interface {
	Subject(context.Context, uuid.UUID) (domain.Subject, error)
	Subjects(context.Context) ([]domain.Subject, error)
	CreateSubject(context.Context, *domain.Subject) error
	UpdateSubject(context.Context, *domain.Subject) error
	DeleteSubject(context.Context, uuid.UUID) error
}

type Service struct {
	repo   domain.SubjectRepository
	logger log.Logger
}

func NewService(repo domain.SubjectRepository, logger log.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

func (s *Service) Subject(_ context.Context, id uuid.UUID) (domain.Subject, error) {
	sub, err := s.repo.Subject(id)
	if err != nil {
		return domain.Subject{}, fmt.Errorf("service can't find subject: %w", err)
	}
	return sub, nil
}

func (s *Service) Subjects(_ context.Context) ([]domain.Subject, error) {
	subs, err := s.repo.Subjects()
	if err != nil {
		return []domain.Subject{}, fmt.Errorf("service didn't found any subject: %w", err)
	}
	return subs, nil
}

func (s *Service) CreateSubject(_ context.Context, sub *domain.Subject) error {
	if err := s.repo.CreateSubject(sub); err != nil {
		return fmt.Errorf("service can't create course: %w", err)
	}
	return nil
}

func (s *Service) UpdateSubject(_ context.Context, sub *domain.Subject) error {
	if err := s.repo.UpdateSubject(sub); err != nil {
		return fmt.Errorf("service can't update course: %w", err)
	}
	return nil
}

func (s *Service) DeleteSubject(_ context.Context, id uuid.UUID) error {
	if err := s.repo.DeleteSubject(id); err != nil {
		return fmt.Errorf("service can't delete course: %w", err)
	}
	return nil
}
