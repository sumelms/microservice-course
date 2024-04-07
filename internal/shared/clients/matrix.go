package clients

import (
	"context"

	"github.com/google/uuid"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

type MatrixClient struct {
	service *domain.Service
}

func NewMatrixClient(svc *domain.Service) *MatrixClient {
	return &MatrixClient{service: svc}
}

func (c MatrixClient) CourseMatrixExists(ctx context.Context, courseUUID uuid.UUID, matrixUUID uuid.UUID) error {
	_, err := c.service.CourseMatrixExists(ctx, courseUUID, matrixUUID)
	if err != nil {
		return err
	}
	return nil
}
