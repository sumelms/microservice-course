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

func (c MatrixClient) MatrixExists(ctx context.Context, matrixUUID uuid.UUID) error {
	_, err := c.service.Matrix(ctx, matrixUUID)
	if err != nil {
		return err
	}
	return nil
}
