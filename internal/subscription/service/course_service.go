package service

import (
	"context"

	"github.com/go-kit/log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/course"
	courseDomain "github.com/sumelms/microservice-course/internal/course/domain"
)

type courseSvc struct {
	course courseDomain.ServiceInterface
}

func NewCourseSvc(db *sqlx.DB, logger log.Logger) (courseSvc, error) { // nolint: revive
	svc, err := course.NewService(db, logger)
	if err != nil {
		return courseSvc{}, err
	}
	return courseSvc{
		course: svc,
	}, nil
}

func (svc courseSvc) ExistCourse(ctx context.Context, id uuid.UUID) error {
	_, err := svc.course.Course(ctx, id)
	return err
}
