package domain

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
)

func (s *Service) Course(_ context.Context, id uuid.UUID) (Course, error) {
	c, err := s.courses.Course(id)
	if err != nil {
		return Course{}, fmt.Errorf("service can't find course: %w", err)
	}
	return c, nil
}

func (s *Service) Courses(_ context.Context) ([]Course, error) {
	cc, err := s.courses.Courses()
	if err != nil {
		return []Course{}, fmt.Errorf("service didn't found any course: %w", err)
	}
	return cc, nil
}

func (s *Service) CreateCourse(_ context.Context, c *Course) error {
	if err := s.courses.CreateCourse(c); err != nil {
		return fmt.Errorf("service can't create course: %w", err)
	}

	// @TODO Dispatch CourseCreatedEvent
	// @TODO Create the Event Pub/Sub Adapter to handle Event Types
	payload, err := json.Marshal(c)
	if err != nil {
		// @TODO What would happen when it fails?
		s.logger.Log("events", "marshal", err)
	}

	msg := message.NewMessage(watermill.NewUUID(), payload)
	err = s.publisher.Publish("course.created", msg)
	if err != nil {
		// @TODO What would happen when it fails?
		s.logger.Log("events", "publisher", err)
	}

	return nil
}

func (s *Service) UpdateCourse(_ context.Context, c *Course) error {
	if err := s.courses.UpdateCourse(c); err != nil {
		return fmt.Errorf("service can't update course: %w", err)
	}
	return nil
}

func (s *Service) DeleteCourse(_ context.Context, id uuid.UUID) error {
	if err := s.courses.DeleteCourse(id); err != nil {
		return fmt.Errorf("service can't delete course: %w", err)
	}
	return nil
}
