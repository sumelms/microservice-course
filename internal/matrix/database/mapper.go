package database

import (
	"github.com/google/uuid"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

func toDBModel(entity *domain.Matrix) Matrix {
	m := Matrix{
		Title:       entity.Title,
		Description: entity.Description,
		CourseID:    uuid.MustParse(entity.CourseID),
	}

	if len(entity.UUID) > 0 {
		m.UUID = uuid.MustParse(entity.UUID)
	}

	if entity.ID > 0 {
		// gorm.Model fields
		m.ID = entity.ID
		m.CreatedAt = entity.CreatedAt
		m.UpdatedAt = entity.UpdatedAt

		if !entity.DeletedAt.IsZero() {
			m.DeletedAt = entity.DeletedAt
		}
	}

	return m
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
		CourseID:    entity.CourseID.String(),
	}
}
