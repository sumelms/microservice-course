package database

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/course/domain"
	utils "github.com/sumelms/microservice-course/tests"
)

var (
	course = domain.Course{
		ID:          1,
		UUID:        utils.CourseUUID,
		Code:        "SUME123",
		Name:        "Course Name",
		Underline:   "Course Underline",
		Image:       "image.png",
		ImageCover:  "image_cover.png",
		Excerpt:     "Course Excerpt",
		Description: "Course Description",
		CreatedAt:   utils.Now,
		UpdatedAt:   utils.Now,
		DeletedAt:   nil,
	}
)

func newCourseTestDB() (*sqlx.DB, sqlmock.Sqlmock, map[string]*sqlmock.ExpectedPrepare) {
	return utils.NewTestDB(queriesCourse())
}

func TestRepository_Course(t *testing.T) {
	validRows := sqlmock.NewRows([]string{"id", "uuid", "code", "name", "underline", "image", "image_cover", "excerpt",
		"description", "created_at", "updated_at", "deleted_at"}).
		AddRow(course.ID, course.UUID, course.Code, course.Name, course.Underline, course.Image, course.ImageCover,
			course.Excerpt, course.Description, course.CreatedAt, course.UpdatedAt, course.DeletedAt)

	type args struct {
		id uuid.UUID
	}

	tests := []struct {
		name    string
		args    args
		rows    *sqlmock.Rows
		want    domain.Course
		wantErr bool
	}{
		{
			name:    "get course",
			args:    args{id: course.UUID},
			rows:    validRows,
			want:    course,
			wantErr: false,
		},
		{
			name:    "course not found error",
			args:    args{id: uuid.MustParse("6cd7a01c-ff18-4cfb-9b35-16e710115c5f")},
			rows:    utils.EmptyRows,
			want:    domain.Course{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newCourseTestDB()
			r, err := NewCourseRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected when creating the courseRepository", err)
			}
			prep, ok := stmts[getCourse]
			if !ok {
				t.Fatalf("prepared statement %s not found", getCourse)
			}

			prep.ExpectQuery().WithArgs(utils.CourseUUID).WillReturnRows(validRows)

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
	validRows := sqlmock.NewRows([]string{"id", "uuid", "code", "name", "underline", "image", "image_cover", "excerpt",
		"description", "created_at", "updated_at", "deleted_at"}).
		AddRow(course.ID, course.UUID, course.Code, course.Name, course.Underline, course.Image, course.ImageCover,
			course.Excerpt, course.Description, course.CreatedAt, course.UpdatedAt, course.DeletedAt).
		AddRow(2, uuid.MustParse("7aec21ad-2fa8-4ddd-b5af-073144031ecc"), course.Code, course.Name,
			course.Underline, course.Image, course.ImageCover, course.Excerpt, course.Description, course.CreatedAt,
			course.UpdatedAt, course.DeletedAt)

	tests := []struct {
		name    string
		rows    *sqlmock.Rows
		wantLen int
		wantErr bool
	}{
		{
			name:    "get all courses",
			rows:    validRows,
			wantLen: 2,
			wantErr: false,
		},
		{
			name:    "get no courses",
			rows:    utils.EmptyRows,
			wantLen: 0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newCourseTestDB()
			r, err := NewCourseRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected when creating the courseRepository", err)
			}
			prep, ok := stmts[listCourse]
			if !ok {
				t.Fatalf("prepared statement %s not found", getCourse)
			}

			prep.ExpectQuery().WillReturnRows(tt.rows)

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
	validRows := sqlmock.NewRows([]string{"id", "uuid", "code", "name", "underline", "image", "image_cover", "excerpt",
		"description", "created_at", "updated_at", "deleted_at"}).
		AddRow(course.ID, course.UUID, course.Code, course.Name, course.Underline, course.Image, course.ImageCover,
			course.Excerpt, course.Description, course.CreatedAt, course.UpdatedAt, course.DeletedAt)

	type args struct {
		c *domain.Course
	}

	tests := []struct {
		name    string
		rows    *sqlmock.Rows
		args    args
		wantErr bool
	}{
		{
			name:    "create course",
			rows:    validRows,
			args:    args{c: &course},
			wantErr: false,
		},
		{
			name:    "empty fields",
			rows:    utils.EmptyRows,
			args:    args{c: &course},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newCourseTestDB()
			r, err := NewCourseRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected when creating the courseRepository", err)
			}
			prep, ok := stmts[createCourse]
			if !ok {
				t.Fatalf("prepared statement %s not found", getCourse)
			}

			prep.ExpectQuery().WillReturnRows(tt.rows)

			if err := r.CreateCourse(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CreateCourse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_UpdateCourse(t *testing.T) {
	validRows := sqlmock.NewRows([]string{"id", "uuid", "code", "name", "underline", "image", "image_cover", "excerpt",
		"description", "created_at", "updated_at", "deleted_at"}).
		AddRow(course.ID, course.UUID, course.Code, course.Name, course.Underline, course.Image, course.ImageCover,
			course.Excerpt, course.Description, course.CreatedAt, course.UpdatedAt, course.DeletedAt)

	type args struct {
		c *domain.Course
	}

	tests := []struct {
		name    string
		args    args
		rows    *sqlmock.Rows
		wantErr bool
	}{
		{
			name:    "update course",
			args:    args{c: &course},
			rows:    validRows,
			wantErr: false,
		},
		{
			name:    "empty course",
			args:    args{c: &domain.Course{}},
			rows:    utils.EmptyRows,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newCourseTestDB()
			r, err := NewCourseRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected when creating the courseRepository", err)
			}
			prep, ok := stmts[updateCourse]
			if !ok {
				t.Fatalf("prepared statement %s not found", updateCourse)
			}

			prep.ExpectQuery().WillReturnRows(tt.rows)

			if err := r.UpdateCourse(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UpdateCourse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
