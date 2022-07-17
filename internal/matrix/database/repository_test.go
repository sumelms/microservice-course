package database

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/matrix/domain"
	"github.com/sumelms/microservice-course/tests/database"
)

var (
	now        = time.Now()
	matrixUUID = uuid.MustParse("dd7c915b-849a-4ba4-bc09-aeecd95c40cc")
	courseUUID = uuid.MustParse("79e1d30d-77f0-4d2f-995c-74aef97c76bf")
	matrix     = domain.Matrix{
		ID:          1,
		UUID:        matrixUUID,
		Code:        "SUME123",
		Name:        "Matrix Name",
		Description: "Matrix Description",
		CourseID:    courseUUID,
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   nil,
	}
	emptyRows = sqlmock.NewRows([]string{})
)

func newTestTB() (*sqlx.DB, sqlmock.Sqlmock, map[string]*sqlmock.ExpectedPrepare) {
	db, mock := database.NewDBMock()

	sqlStatements := make(map[string]*sqlmock.ExpectedPrepare)
	for queryName, query := range queries() {
		stmt := mock.ExpectPrepare(fmt.Sprintf("^%s$", regexp.QuoteMeta(string(query))))
		sqlStatements[queryName] = stmt
	}

	mock.MatchExpectationsInOrder(false)
	return db, mock, sqlStatements
}

func TestRepository_Matrix(t *testing.T) {
	validRows := sqlmock.NewRows([]string{"id", "uuid", "code", "name", "description",
		"course_id", "created_at", "updated_at", "deleted_at"}).
		AddRow(matrix.ID, matrix.UUID, matrix.Code, matrix.Name, matrix.Description,
			matrix.CourseID, matrix.CreatedAt, matrix.UpdatedAt, matrix.DeletedAt)

	type args struct {
		id uuid.UUID
	}

	tests := []struct {
		name    string
		args    args
		rows    *sqlmock.Rows
		want    domain.Matrix
		wantErr bool
	}{
		{
			name:    "get matrix",
			args:    args{id: matrixUUID},
			rows:    validRows,
			want:    matrix,
			wantErr: false,
		},
		{
			name:    "matrix not found error",
			args:    args{id: uuid.MustParse("8281f61e-956e-4f64-ac0e-860c444c5f86")},
			rows:    emptyRows,
			want:    domain.Matrix{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newTestTB()
			r, err := NewRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected when creating the repository", err)
			}
			prep, ok := stmts[getMatrix]
			if !ok {
				t.Fatalf("prepared statement %s not found", getMatrix)
			}

			prep.ExpectQuery().WithArgs(matrixUUID).WillReturnRows(tt.rows)

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
	validRows := sqlmock.NewRows([]string{"id", "uuid", "code", "name", "description",
		"course_id", "created_at", "updated_at", "deleted_at"}).
		AddRow(matrix.ID, matrix.UUID, matrix.Code, matrix.Name, matrix.Description,
			matrix.CourseID, matrix.CreatedAt, matrix.UpdatedAt, matrix.DeletedAt).
		AddRow(2, uuid.MustParse("e74868b2-72d4-4591-a90d-122a9ac2d945"),
			matrix.Code, matrix.Name, matrix.Description, matrix.CourseID, matrix.CreatedAt,
			matrix.UpdatedAt, matrix.DeletedAt)

	tests := []struct {
		name    string
		rows    *sqlmock.Rows
		wantLen int
		wantErr bool
	}{
		{
			name:    "get all matrices",
			rows:    validRows,
			wantLen: 2,
			wantErr: false,
		},
		{
			name:    "get no matrices",
			rows:    emptyRows,
			wantLen: 0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newTestTB()
			r, err := NewRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected when creating the repository", err)
			}
			prep, ok := stmts[listMatrix]
			if !ok {
				t.Fatalf("prepared statement %s not found", listMatrix)
			}

			prep.ExpectQuery().WillReturnRows(tt.rows)

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
