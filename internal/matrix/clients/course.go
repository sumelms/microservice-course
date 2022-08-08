package clients

import (
	"context"

	"github.com/google/uuid"

	"github.com/sumelms/microservice-course/internal/course/domain"
)

type courseClient struct {
	service *domain.Service
}

func NewCourseClient(svc *domain.Service) *courseClient {
	return &courseClient{service: svc}
}

func (c courseClient) CourseExists(ctx context.Context, id uuid.UUID) error {
	_, err := c.service.Course(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
