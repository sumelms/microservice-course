package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) Course(_ context.Context, courseUUID uuid.UUID) (Course, error) {
	course, err := s.courses.Course(courseUUID)
	if err != nil {
		return Course{}, fmt.Errorf("service can't find course: %w", err)
	}

	return course, nil
}

func (s *Service) Courses(_ context.Context) ([]Course, error) {
	cc, err := s.courses.Courses()
	if err != nil {
		return []Course{}, fmt.Errorf("service didn't found any course: %w", err)
	}

	return cc, nil
}

func (s *Service) CreateCourse(_ context.Context, course *Course) error {
	if err := s.courses.CreateCourse(course); err != nil {
		return fmt.Errorf("service can't create course: %w", err)
	}

	return nil
}

func (s *Service) UpdateCourse(_ context.Context, course *Course) error {
	if err := s.courses.UpdateCourse(course); err != nil {
		return fmt.Errorf("service can't update course: %w", err)
	}
	return nil
}

func (s *Service) DeleteCourse(_ context.Context, course *DeletedCourse) error {
	if err := s.courses.DeleteCourse(course); err != nil {
		return fmt.Errorf("service can't delete course: %w", err)
	}

	return nil
}
