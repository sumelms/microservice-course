package database

import (
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/tests"
)

var now = time.Now()

func newTestCourse() domain.Course {
	return domain.Course{
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
}

func TestRepository_Course(t *testing.T) {
	db, mock := tests.NewDBMock()

	c := newTestCourse()
	rows := mock.NewRows([]string{"id", "uuid", "title", "subtitle", "excerpt",
		"description", "created_at", "updated_at", "deleted_at"}).
		AddRow(c.ID, c.UUID, c.Title, c.Subtitle, c.Excerpt,
			c.Description, c.CreatedAt, c.UpdatedAt, c.DeletedAt)

	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		id uuid.UUID
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		rows    *sqlmock.Rows
		want    domain.Course
		wantErr bool
	}{
		{
			name:    "get course",
			fields:  fields{DB: db},
			args:    args{id: c.UUID},
			rows:    rows,
			want:    c,
			wantErr: false,
		},
		{
			name:    "course not found error",
			fields:  fields{DB: db},
			args:    args{id: uuid.MustParse("8281f61e-956e-4f64-ac0e-860c444c5f86")},
			rows:    rows,
			want:    domain.Course{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := &Repository{DB: tt.fields.DB}
			defer func() {
				_ = r.Close()
			}()

			query := "SELECT \\* FROM courses WHERE deleted_at IS NULL AND uuid = \\$1"
			mock.ExpectQuery(query).WithArgs(tt.args.id).WillReturnRows(tt.rows)

			got, err := r.Course(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Course() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Course() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_Courses(t *testing.T) {
	db, mock := tests.NewDBMock()
	c := newTestCourse()

	rows := mock.NewRows([]string{"id", "uuid", "title", "subtitle", "excerpt", "description",
		"created_at", "updated_at", "deleted_at"}).
		AddRow(c.ID, c.UUID, c.Title, c.Subtitle, c.Excerpt, c.Description,
			c.CreatedAt, c.UpdatedAt, c.DeletedAt).
		AddRow(2, uuid.MustParse("7aec21ad-2fa8-4ddd-b5af-073144031ecc"), c.Title, c.Subtitle, c.Excerpt, c.Description,
			c.CreatedAt, c.UpdatedAt, c.DeletedAt)

	type fields struct {
		DB *sqlx.DB
	}

	tests := []struct {
		name    string
		fields  fields
		rows    *sqlmock.Rows
		wantLen int
		wantErr bool
	}{
		{
			name:    "get all courses",
			fields:  fields{DB: db},
			rows:    rows,
			wantLen: 2,
			wantErr: false,
		},
		{
			name:    "get no courses",
			fields:  fields{DB: db},
			rows:    nil,
			wantLen: 0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := &Repository{DB: tt.fields.DB}
			defer func() {
				_ = r.Close()
			}()

			query := "SELECT \\* FROM courses WHERE deleted_at IS NULL"
			mock.ExpectQuery(query).WillReturnRows(rows)

			got, err := r.Courses()
			if (err != nil) != tt.wantErr {
				t.Errorf("Courses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("Courses() got = %v, want %v", got, tt.wantLen)
			}
		})
	}
}

func TestRepository_CreateCourse(t *testing.T) {
	db, mock := tests.NewDBMock()

	c := newTestCourse()
	rows := mock.NewRows([]string{"id", "uuid", "title", "subtitle", "excerpt", "description",
		"created_at", "updated_at", "deleted_at"}).
		AddRow(c.ID, c.UUID, c.Title, c.Subtitle, c.Excerpt, c.Description,
			c.CreatedAt, c.UpdatedAt, c.DeletedAt)

	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		c *domain.Course
	}

	tests := []struct {
		name    string
		fields  fields
		rows    *sqlmock.Rows
		args    args
		wantErr bool
	}{
		{
			name:    "create course",
			fields:  fields{DB: db},
			rows:    rows,
			args:    args{c: &c},
			wantErr: false,
		},
		{
			name:    "empty fields",
			fields:  fields{DB: db},
			rows:    nil,
			args:    args{c: &c},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := &Repository{DB: tt.fields.DB}
			defer func() {
				_ = r.Close()
			}()

			query := "INSERT INTO courses \\(title, subtitle, excerpt, description\\) VALUES \\(\\$1, \\$2, \\$3, \\$4\\) RETURNING \\*"
			mock.ExpectQuery(query).WithArgs(c.Title, c.Subtitle, c.Excerpt, c.Description).WillReturnRows(rows)

			if err := r.CreateCourse(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CreateCourse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_UpdateCourse(t *testing.T) {
	db, mock := tests.NewDBMock()

	c := newTestCourse()
	rows := mock.NewRows([]string{"id", "uuid", "title", "subtitle", "excerpt", "description",
		"created_at", "updated_at", "deleted_at"}).
		AddRow(c.ID, c.UUID, c.Title, c.Subtitle, c.Excerpt, c.Description,
			c.CreatedAt, c.UpdatedAt, c.DeletedAt)

	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		c *domain.Course
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		rows    *sqlmock.Rows
		wantErr bool
	}{
		{
			name:    "update course",
			fields:  fields{DB: db},
			args:    args{c: &c},
			rows:    rows,
			wantErr: false,
		},
		{
			name:    "empty course",
			fields:  fields{DB: db},
			args:    args{c: &domain.Course{}},
			rows:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{DB: tt.fields.DB}
			defer func() {
				_ = r.Close()
			}()

			query := "UPDATE courses SET title = \\$1, subtitle = \\$2, excerpt = \\$3, description = \\$4 WHERE uuid = \\$5 RETURNING \\*"
			mock.ExpectQuery(query).WithArgs(c.Title, c.Subtitle, c.Excerpt, c.Description, c.UUID).WillReturnRows(rows)

			if err := r.UpdateCourse(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UpdateCourse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
