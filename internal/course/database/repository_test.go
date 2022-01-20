package database

import (
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

	"github.com/sumelms/microservice-course/internal/course/domain"
)

var now = time.Now()

var c = &domain.Course{
	ID:          1,
	UUID:        uuid.MustParse("dd7c915b-849a-4ba4-bc09-aeecd95c40cc"),
	Title:       "Course Title",
	Subtitle:    "Course Subtitle",
	Excerpt:     "Course Excerpt",
	Description: "Course Description",
	CreatedAt:   now,
	UpdatedAt:   now,
	DeletedAt:   nil,
}

func NewDBMock() (*sqlx.DB, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db := sqlx.NewDb(mockDB, "sqlmock")
	return db, mock
}

func TestRepository_Course(t *testing.T) {
	db, mock := NewDBMock()

	repo := &Repository{db}
	defer func() {
		_ = repo.Close()
	}()

	query := "SELECT \\* FROM courses WHERE uuid = \\$1"
	rows := sqlmock.
		NewRows([]string{"id", "uuid", "title", "subtitle", "excerpt", "description",
			"created_at", "updated_at", "deleted_at"}).
		AddRow(c.ID, c.UUID, c.Title, c.Subtitle, c.Excerpt, c.Description,
			c.CreatedAt, c.UpdatedAt, c.DeletedAt)

	mock.ExpectQuery(query).WithArgs(c.UUID).WillReturnRows(rows)

	course, err := repo.Course(c.UUID)
	assert.NotNil(t, course)
	assert.NoError(t, err)
}

func TestRepository_Courses(t *testing.T) {
	db, mock := NewDBMock()

	repo := &Repository{db}
	defer func() {
		_ = repo.Close()
	}()

	query := "SELECT \\* FROM courses"
	rows := sqlmock.
		NewRows([]string{"id", "uuid", "title", "subtitle", "excerpt", "description",
			"created_at", "updated_at", "deleted_at"}).
		AddRow(c.ID, c.UUID, c.Title, c.Subtitle, c.Excerpt, c.Description,
			c.CreatedAt, c.UpdatedAt, c.DeletedAt).
		AddRow(2, uuid.MustParse("7aec21ad-2fa8-4ddd-b5af-073144031ecc"), c.Title, c.Subtitle, c.Excerpt, c.Description,
			c.CreatedAt, c.UpdatedAt, c.DeletedAt)

	mock.ExpectQuery(query).WillReturnRows(rows)

	courses, err := repo.Courses()
	assert.NotNil(t, courses)
	assert.Len(t, courses, 2)
	assert.NoError(t, err)
}
