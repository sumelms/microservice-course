package database

import (
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/matrix/domain"
	mtests "github.com/sumelms/microservice-course/tests/database"
)

var now = time.Now()

func newTestMatrix() domain.Matrix {
	return domain.Matrix{
		ID:          1,
		UUID:        uuid.MustParse("dd7c915b-849a-4ba4-bc09-aeecd95c40cc"),
		Title:       "Matrix Name",
		Description: "Matrix Description",
		CourseID:    uuid.MustParse("79e1d30d-77f0-4d2f-995c-74aef97c76bf"),
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   nil,
	}
}

func TestRepository_Matrix(t *testing.T) {
	db, mock := mtests.NewDBMock()

	m := newTestMatrix()
	rows := mock.NewRows([]string{"id", "uuid", "title", "description",
		"course_id", "created_at", "updated_at", "deleted_at"}).
		AddRow(m.ID, m.UUID, m.Title, m.Description,
			m.CourseID, m.CreatedAt, m.UpdatedAt, m.DeletedAt)

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
		want    domain.Matrix
		wantErr bool
	}{
		{
			name:    "get matrix",
			fields:  fields{DB: db},
			args:    args{id: m.UUID},
			rows:    rows,
			want:    m,
			wantErr: false,
		},
		{
			name:    "matrix not found error",
			fields:  fields{DB: db},
			args:    args{id: uuid.MustParse("8281f61e-956e-4f64-ac0e-860c444c5f86")},
			rows:    rows,
			want:    domain.Matrix{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := &repository{DB: tt.fields.DB}
			defer func() {
				_ = r.Close()
			}()

			query := "SELECT \\* FROM matrices WHERE deleted_at IS NULL AND uuid = \\$1"
			mock.ExpectQuery(query).WithArgs(tt.args.id).WillReturnRows(tt.rows)

			got, err := r.Matrix(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_Matrices(t *testing.T) {
	db, mock := mtests.NewDBMock()

	m := newTestMatrix()
	rows := mock.NewRows([]string{"id", "uuid", "title", "description",
		"course_id", "created_at", "updated_at", "deleted_at"}).
		AddRow(m.ID, m.UUID, m.Title, m.Description,
			m.CourseID, m.CreatedAt, m.UpdatedAt, m.DeletedAt).
		AddRow(2, uuid.MustParse("e74868b2-72d4-4591-a90d-122a9ac2d945"), m.Title, m.Description,
			m.CourseID, m.CreatedAt, m.UpdatedAt, m.DeletedAt)

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
			name:    "get all matrices",
			fields:  fields{DB: db},
			rows:    rows,
			wantLen: 2,
			wantErr: false,
		},
		{
			name:    "get no matrices",
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

			r := &repository{DB: tt.fields.DB}
			defer func() {
				_ = r.Close()
			}()

			query := "SELECT \\* FROM matrices WHERE deleted_at IS NULL"
			mock.ExpectQuery(query).WillReturnRows(rows)

			got, err := r.Matrices()
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("Matrices() got = %v, want %v", got, tt.wantLen)
			}
		})
	}
}
