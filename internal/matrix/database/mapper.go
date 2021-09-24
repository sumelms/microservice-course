package database

import (
	"github.com/google/uuid"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

func toDBModel(entity *domain.Matrix) Matrix {
	matrix := Matrix{
		Title:       entity.Title,
		Description: entity.Description,
		CourseID:    entity.CourseID,
	}

	if len(entity.UUID) > 0 {
		matrix.UUID = uuid.MustParse(entity.UUID)
	}

	if entity.ID > 0 {
		// gorm.Model fields
		matrix.ID = entity.ID
		matrix.CreatedAt = entity.CreatedAt
		matrix.UpdatedAt = entity.UpdatedAt

		if !entity.DeletedAt.IsZero() {
			matrix.DeletedAt = entity.DeletedAt
		}
	}

	return matrix
}

func toDomainModel(entity *Matrix) domain.Matrix {
	return domain.Matrix{
		ID:          entity.ID,
		UUID:        entity.UUID.String(),
		Title:       entity.Title,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
		CourseID:    entity.CourseID,
	}
}
