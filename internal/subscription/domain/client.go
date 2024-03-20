package domain

import (
	"context"

	"github.com/google/uuid"
)

type CourseClient interface {
	CourseExists(ctx context.Context, courseUUID uuid.UUID) error
}

type MatrixClient interface {
	CourseMatrixExists(ctx context.Context, courseUUID uuid.UUID, matrixUUID uuid.UUID) error
}
