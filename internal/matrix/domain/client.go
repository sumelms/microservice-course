package domain

import (
	"context"

	"github.com/google/uuid"
)

type CourseClient interface {
	CourseExists(ctx context.Context, id uuid.UUID) error
}
