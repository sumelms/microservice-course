package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) Subject(_ context.Context, subjectUUID uuid.UUID) (Subject, error) {
	subject, err := s.subjects.Subject(subjectUUID)
	if err != nil {
		return Subject{}, fmt.Errorf("service can't find subject: %w", err)
	}
	return subject, nil
}

func (s *Service) Subjects(_ context.Context) ([]Subject, error) {
	subjects, err := s.subjects.Subjects()
	if err != nil {
		return []Subject{}, fmt.Errorf("service didn't found any subject: %w", err)
	}
	return subjects, nil
}

func (s *Service) CreateSubject(_ context.Context, subject *Subject) error {
	if err := s.subjects.CreateSubject(subject); err != nil {
		return fmt.Errorf("service can't create subject: %w", err)
	}
	return nil
}

func (s *Service) UpdateSubject(_ context.Context, subject *Subject) error {
	if err := s.subjects.UpdateSubject(subject); err != nil {
		return fmt.Errorf("service can't update subject: %w", err)
	}
	return nil
}

func (s *Service) DeleteSubject(_ context.Context, subject *DeletedSubject) error {
	if err := s.subjects.DeleteSubject(subject); err != nil {
		return fmt.Errorf("service can't delete subject: %w", err)
	}
	return nil
}
