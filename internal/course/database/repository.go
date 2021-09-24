package database

import (
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-course/internal/course/domain"
	merrors "github.com/sumelms/microservice-course/pkg/errors"
)

const (
	whereCourseUUID = "uuid = ?"
)

// Repository struct
type Repository struct {
	db     *gorm.DB
	logger log.Logger
}

// NewRepository creates a new profile repository
func NewRepository(db *gorm.DB, logger log.Logger) *Repository {
	db.AutoMigrate(&Course{})

	return &Repository{
		db:     db,
		logger: logger,
	}
}

// List courses
func (r *Repository) List() ([]domain.Course, error) {
	var courses []Course

	query := r.db.Find(&courses)
	if query.RecordNotFound() {
		return []domain.Course{}, nil
	}
	if err := query.Error; err != nil {
		return []domain.Course{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "list courses")
	}

	list := make([]domain.Course, len(courses))
	for i := range courses {
		course := courses[i]
		list = append(list, toDomainModel(&course))
	}
	return list, nil
}

// Create creates a course
func (r *Repository) Create(course *domain.Course) (domain.Course, error) {
	entity := toDBModel(course)

	if err := r.db.Create(&entity).Error; err != nil {
		return domain.Course{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "can't create course")
	}
	return toDomainModel(&entity), nil
}

// Find get a course by its ID
func (r *Repository) Find(id string) (domain.Course, error) {
	var course Course

	query := r.db.Where(whereCourseUUID, id).First(&course)
	if query.RecordNotFound() {
		return domain.Course{}, merrors.NewErrorf(merrors.ErrCodeNotFound, "course not found")
	}
	if err := query.Error; err != nil {
		return domain.Course{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "find course")
	}

	return toDomainModel(&course), nil
}

// Update the given course
func (r *Repository) Update(c *domain.Course) (domain.Course, error) {
	var course Course

	query := r.db.Where(whereCourseUUID, c.UUID).First(&course)

	if query.RecordNotFound() {
		return domain.Course{}, merrors.NewErrorf(merrors.ErrCodeNotFound, "course not found")
	}

	c.ID = course.ID
	query = r.db.Save(&c)

	if err := query.Error; err != nil {
		return domain.Course{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "can't update course")
	}

	return *c, nil
}

// Delete a course by its ID
func (r *Repository) Delete(id string) error {
	query := r.db.Where(whereCourseUUID, id).Delete(&Course{})

	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return merrors.WrapErrorf(err, merrors.ErrCodeNotFound, "course not found")
		}
		return merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "delete course")
	}

	return nil
}
