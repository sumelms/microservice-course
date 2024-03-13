package domain

import (
	"context"

	"github.com/google/uuid"
)

type CourseClient interface {
	CourseExists(ctx context.Context, courseUUID uuid.UUID) error
}

type MatrixClient interface {
	MatrixExists(ctx context.Context, matrixUUID uuid.UUID) error
}
