package database

import (
	"github.com/sumelms/microservice-course/internal/course/domain"
)

func toDBModel(entity *domain.Course) Course {
	course := Course{}

	if entity.DeletedAt != nil {
		course.DeletedAt = entity.DeletedAt
	}

	return course
}

func toDomainModel(entity *Course) domain.Course {
	return domain.Course{}
}
