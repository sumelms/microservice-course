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

	if entity.DeletedAt != nil {
		course.DeletedAt = entity.DeletedAt
	}

	return course
}

func toDomainModel(entity *Course) domain.Course {
	course := domain.Course{
		UUID:        entity.UUID.String(),
		Title:       entity.Title,
		Subtitle:    entity.Subtitle,
		Excerpt:     entity.Excerpt,
		Description: entity.Description,
	}

	return course
}
