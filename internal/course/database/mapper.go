package database

import (
	"github.com/google/uuid"
	"github.com/sumelms/microservice-course/internal/course/domain"
)

func toDBModel(entity *domain.Course) Course {
	course := Course{
		Title:       entity.Title,
		Subtitle:    entity.Subtitle,
		Excerpt:     entity.Excerpt,
		Description: entity.Description,
	}

	if len(entity.UUID) > 0 {
		course.UUID = uuid.MustParse(entity.UUID)
	}

	if entity.ID > 0 {
		// gorm.Model fields
		course.ID = entity.ID
		course.CreatedAt = entity.CreatedAt
		course.UpdatedAt = entity.UpdatedAt

		if !entity.DeletedAt.IsZero() {
			course.DeletedAt = entity.DeletedAt
		}
	}
	return course
}

func toDomainModel(entity *Course) domain.Course {
	return domain.Course{
		ID:          entity.ID,
		UUID:        entity.UUID.String(),
		Title:       entity.Title,
		Subtitle:    entity.Subtitle,
		Excerpt:     entity.Excerpt,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
	}
}
