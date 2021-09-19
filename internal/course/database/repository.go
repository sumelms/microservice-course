package database

import (
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-course/internal/course/domain"
	merrors "github.com/sumelms/microservice-course/pkg/errors"
)

const (
	whereCourseID = "id = ?"
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

// Create creates a course
func (r *Repository) Create(course *domain.Course) (domain.Course, error) {
	entity := toDBModel(course)

	if err := r.db.Create(entity).Error; err != nil {
		return domain.Course{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "create course")
	}
	return toDomainModel(&entity), nil
}

//// Find get a course by its ID
//func (r *Repository) Find(courseID domain.UserID) (domain.Course, error) {
//	var course Course
//
//	query := r.db.Where(whereCourseID, courseID).First(&course)
//
//	if query.RecordNotFound() {
//		return domain.Course{}, merrors.NewErrorf(merrors.ErrCodeNotFound, "course not found")
//	}
//	if err := query.Error; err != nil {
//		return domain.Course{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "find course")
//	}
//
//	return toDomainModel(&course), nil
//}
//
//// Update the given course
//func (r *Repository) Update(p *domain.Course) (domain.Course, error) {
//	var course Course
//
//	query := r.db.Where(whereCourseID, p.ID).First(&course)
//
//	if query.RecordNotFound() {
//		return domain.Course{}, merrors.NewErrorf(merrors.ErrCodeNotFound, "course not found")
//	}
//
//	query = r.db.Save(&course)
//
//	if err := query.Error; err != nil {
//		return domain.Course{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "update course")
//	}
//
//	return toDomainModel(&course), nil
//}
//
//// Delete a profile by CourseID
//func (r *Repository) Delete(id domain.UserID) error {
//	query := r.db.Where(whereCourseID, id).Delete(&Course{})
//
//	if err := query.Error; err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return merrors.WrapErrorf(err, merrors.ErrCodeNotFound, "course not found")
//		}
//		return merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "delete course")
//	}
//
//	return nil
//}
