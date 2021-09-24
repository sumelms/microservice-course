package database

import (
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
	merrors "github.com/sumelms/microservice-course/pkg/errors"
)

const (
	whereMatrixUUID = "UUID = ?"
)

type Repository struct {
	db     *gorm.DB
	logger log.Logger
}

func NewRepository(db *gorm.DB, logger log.Logger) *Repository {
	db.AutoMigrate(&Matrix{})

	return &Repository{
		db:     db,
		logger: logger,
	}
}

func (r *Repository) Create(matrix *domain.Matrix) (domain.Matrix, error) {
	entity := toDBModel(matrix)

	if err := r.db.Create(&entity).Error; err != nil {
		return domain.Matrix{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "can't create matrix")
	}
	return toDomainModel(&entity), nil
}

func (r *Repository) Find(id string) (domain.Matrix, error) {
	var matrix Matrix

	query := r.db.Where(whereMatrixUUID, id).First(&matrix)
	if query.RecordNotFound() {
		return domain.Matrix{}, merrors.NewErrorf(merrors.ErrCodeNotFound, "matrix not found")
	}
	if err := query.Error; err != nil {
		return domain.Matrix{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "find matrix")
	}

	return toDomainModel(&matrix), nil
}

func (r *Repository) Update(m *domain.Matrix) (domain.Matrix, error) {
	var matrix Matrix

	query := r.db.Where(whereMatrixUUID, m.UUID).First(&matrix)

	if query.RecordNotFound() {
		return domain.Matrix{}, merrors.NewErrorf(merrors.ErrCodeNotFound, "matrix not found")
	}

	query = r.db.Model(&matrix).Updates(&m)

	if err := query.Error; err != nil {
		return domain.Matrix{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "can't update matrix")
	}

	return *m, nil
}

func (r *Repository) Delete(id string) error {
	query := r.db.Where(whereMatrixUUID, id).Delete(&Matrix{})

	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return merrors.WrapErrorf(err, merrors.ErrCodeNotFound, "matrix not found")
		}
	}

	return nil
}

func (r *Repository) List() ([]domain.Matrix, error) {
	var matrices []Matrix

	query := r.db.Find(&matrices)
	if query.RecordNotFound() {
		return []domain.Matrix{}, nil
	}
	if err := query.Error; err != nil {
		return []domain.Matrix{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "list matrices")
	}

	var list []domain.Matrix
	for _, m := range matrices {
		list = append(list, toDomainModel(&m))
	}

	return list, nil
}
