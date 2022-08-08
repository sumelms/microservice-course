package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) Subject(_ context.Context, id uuid.UUID) (Subject, error) {
	sub, err := s.subjects.Subject(id)
	if err != nil {
		return Subject{}, fmt.Errorf("service can't find subject: %w", err)
	}
	return sub, nil
}

func (s *Service) Subjects(_ context.Context) ([]Subject, error) {
	subs, err := s.subjects.Subjects()
	if err != nil {
		return []Subject{}, fmt.Errorf("service didn't found any subject: %w", err)
	}
	return subs, nil
}

func (s *Service) CreateSubject(_ context.Context, sub *Subject) error {
	if err := s.subjects.CreateSubject(sub); err != nil {
		return fmt.Errorf("service can't create course: %w", err)
	}
	return nil
}

func (s *Service) UpdateSubject(_ context.Context, sub *Subject) error {
	if err := s.subjects.UpdateSubject(sub); err != nil {
		return fmt.Errorf("service can't update course: %w", err)
	}
	return nil
}

func (s *Service) DeleteSubject(_ context.Context, id uuid.UUID) error {
	if err := s.subjects.DeleteSubject(id); err != nil {
		return fmt.Errorf("service can't delete course: %w", err)
	}
	return nil
}
