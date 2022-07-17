package database

import (
	"log"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/course/domain"
	database "github.com/sumelms/microservice-course/tests/database"
)

var (
	now        = time.Now()
	courseUUID = uuid.MustParse("dd7c915b-849a-4ba4-bc09-aeecd95c40cc")
	course     = domain.Course{
		ID:          1,
		UUID:        courseUUID,
		Code:        "SUME123",
		Name:        "Course Name",
		Underline:   "Course Underline",
		Image:       "image.png",
		ImageCover:  "image_cover.png",
		Excerpt:     "Course Excerpt",
		Description: "Course Description",
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   nil,
	}
)

func newTestDB() (*sqlx.DB, sqlmock.Sqlmock, map[string]*sqlmock.ExpectedPrepare) {
	db, mock := database.NewDBMock()

	sqlStatements := make(map[string]*sqlmock.ExpectedPrepare)
	for queryName, query := range queries() {
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(string(query)))
		sqlStatements[queryName] = stmt
	}

	mock.MatchExpectationsInOrder(false)
	return db, mock, sqlStatements
}

func TestRepository_Course(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Course
		wantErr bool
	}{
		{
			name:    "get course",
			args:    args{id: course.UUID},
			want:    course,
			wantErr: false,
		},
		{
			name:    "course not found error",
			args:    args{id: uuid.MustParse("6cd7a01c-ff18-4cfb-9b35-16e710115c5f")},
			want:    domain.Course{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, mock, stmts := newTestDB()
			defer db.Close()

			r, err := NewRepository(db)
			if err != nil {
				log.Fatalf("an error '%s' was not expected when creating the repository", err)
			}

			prep, ok := stmts[getCourse]
			if !ok {
				log.Fatalf("prepared statement %s not found", string(getCourse))
			}

			validRows := mock.NewRows([]string{"id", "uuid", "code", "name", "underline", "image", "image_cover", "excerpt",
				"description", "created_at", "updated_at", "deleted_at"}).
				AddRow(course.ID, course.UUID, course.Code, course.Name, course.Underline, course.Image, course.ImageCover,
					course.Excerpt, course.Description, course.CreatedAt, course.UpdatedAt, course.DeletedAt)

			prep.ExpectQuery().WithArgs(courseUUID).WillReturnRows(validRows)

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

//
//func TestRepository_Courses(t *testing.T) {
//	db, mock := mtests.NewDBMock()
//	c := newTestCourse()
//
//	rows := mock.NewRows([]string{"id", "uuid", "title", "subtitle", "excerpt", "description",
//		"created_at", "updated_at", "deleted_at"}).
//		AddRow(c.ID, c.UUID, c.Name, c.Underline, c.Excerpt, c.Description,
//			c.CreatedAt, c.UpdatedAt, c.DeletedAt).
//		AddRow(2, uuid.MustParse("7aec21ad-2fa8-4ddd-b5af-073144031ecc"), c.Name, c.Underline, c.Excerpt, c.Description,
//			c.CreatedAt, c.UpdatedAt, c.DeletedAt)
//
//	type fields struct {
//		DB *sqlx.DB
//	}
//
//	tests := []struct {
//		name    string
//		fields  fields
//		rows    *sqlmock.Rows
//		wantLen int
//		wantErr bool
//	}{
//		{
//			name:    "get all courses",
//			fields:  fields{DB: db},
//			rows:    rows,
//			wantLen: 2,
//			wantErr: false,
//		},
//		{
//			name:    "get no courses",
//			fields:  fields{DB: db},
//			rows:    nil,
//			wantLen: 0,
//			wantErr: false,
//		},
//	}
//
//	for _, tt := range tests {
//		tt := tt
//		t.Run(tt.name, func(t *testing.T) {
//			t.Parallel()
//
//			r := &Repository{DB: tt.fields.DB}
//
//			query := "SELECT \\* FROM courses WHERE deleted_at IS NULL"
//			mock.ExpectQuery(query).WillReturnRows(rows)
//
//			got, err := r.Courses()
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Courses() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if len(got) != tt.wantLen {
//				t.Errorf("Courses() got = %v, want %v", got, tt.wantLen)
//			}
//		})
//	}
//}
//
//func TestRepository_CreateCourse(t *testing.T) {
//	db, mock := mtests.NewDBMock()
//
//	c := newTestCourse()
//	rows := mock.NewRows([]string{"id", "uuid", "title", "subtitle", "excerpt", "description",
//		"created_at", "updated_at", "deleted_at"}).
//		AddRow(c.ID, c.UUID, c.Name, c.Underline, c.Excerpt, c.Description,
//			c.CreatedAt, c.UpdatedAt, c.DeletedAt)
//
//	type fields struct {
//		DB *sqlx.DB
//	}
//	type args struct {
//		c *domain.Course
//	}
//
//	tests := []struct {
//		name    string
//		fields  fields
//		rows    *sqlmock.Rows
//		args    args
//		wantErr bool
//	}{
//		{
//			name:    "create course",
//			fields:  fields{DB: db},
//			rows:    rows,
//			args:    args{c: &c},
//			wantErr: false,
//		},
//		{
//			name:    "empty fields",
//			fields:  fields{DB: db},
//			rows:    nil,
//			args:    args{c: &c},
//			wantErr: true,
//		},
//	}
//
//	for _, tt := range tests {
//		tt := tt
//		t.Run(tt.name, func(t *testing.T) {
//			t.Parallel()
//
//			r := &Repository{DB: tt.fields.DB}
//
//			query := "INSERT INTO courses \\(title, subtitle, excerpt, description\\) VALUES \\(\\$1, \\$2, \\$3, \\$4\\) RETURNING \\*"
//			mock.ExpectQuery(query).WithArgs(c.Name, c.Underline, c.Excerpt, c.Description).WillReturnRows(rows)
//
//			if err := r.CreateCourse(tt.args.c); (err != nil) != tt.wantErr {
//				t.Errorf("CreateCourse() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestRepository_UpdateCourse(t *testing.T) {
//	db, mock := mtests.NewDBMock()
//
//	c := newTestCourse()
//	rows := mock.NewRows([]string{"id", "uuid", "title", "subtitle", "excerpt", "description",
//		"created_at", "updated_at", "deleted_at"}).
//		AddRow(c.ID, c.UUID, c.Name, c.Underline, c.Excerpt, c.Description,
//			c.CreatedAt, c.UpdatedAt, c.DeletedAt)
//
//	type fields struct {
//		DB *sqlx.DB
//	}
//	type args struct {
//		c *domain.Course
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		rows    *sqlmock.Rows
//		wantErr bool
//	}{
//		{
//			name:    "update course",
//			fields:  fields{DB: db},
//			args:    args{c: &c},
//			rows:    rows,
//			wantErr: false,
//		},
//		{
//			name:    "empty course",
//			fields:  fields{DB: db},
//			args:    args{c: &domain.Course{}},
//			rows:    nil,
//			wantErr: true,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			t.Parallel()
//
//			r := &Repository{DB: tt.fields.DB}
//
//			query := "UPDATE courses SET title = \\$1, subtitle = \\$2, excerpt = \\$3, description = \\$4 WHERE uuid = \\$5 RETURNING \\*"
//			mock.ExpectQuery(query).WithArgs(c.Name, c.Underline, c.Excerpt, c.Description, c.UUID).WillReturnRows(rows)
//
//			if err := r.UpdateCourse(tt.args.c); (err != nil) != tt.wantErr {
//				t.Errorf("UpdateCourse() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
