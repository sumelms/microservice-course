package database

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/matrix/domain"
	utils "github.com/sumelms/microservice-course/tests"
)

var (
	matrix = domain.Matrix{
		UUID:        utils.MatrixUUID,
		CourseUUID:  utils.CourseUUID,
		Code:        "Code",
		Name:        "Name",
		Description: "Description",
		CreatedAt:   utils.Now,
		UpdatedAt:   utils.Now,
	}
)

func newMatrixTestDB() (*sqlx.DB, sqlmock.Sqlmock, map[string]*sqlmock.ExpectedPrepare) {
	return utils.NewTestDB(queriesMatrix())
}

func TestRepository_Matrix(t *testing.T) {
	validRows := sqlmock.NewRows([]string{
		"uuid", "code", "name", "description", "course_uuid",
		"created_at", "updated_at"}).
		AddRow(
			matrix.UUID, matrix.Code, matrix.Name, matrix.Description, matrix.CourseUUID,
			matrix.CreatedAt, matrix.UpdatedAt)

	type args struct {
		UUID uuid.UUID
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
			args:    args{UUID: utils.MatrixUUID},
			rows:    validRows,
			want:    matrix,
			wantErr: false,
		},
		{
			name:    "matrix not found error",
			args:    args{UUID: uuid.MustParse("8281f61e-956e-4f64-ac0e-860c444c5f86")},
			rows:    utils.EmptyRows,
			want:    domain.Matrix{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newMatrixTestDB()
			r, err := NewMatrixRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected when creating the MatrixRepository", err)
			}
			prep, ok := stmts[getMatrix]
			if !ok {
				t.Fatalf("prepared statement %s not found", getMatrix)
			}

			prep.ExpectQuery().WithArgs(utils.MatrixUUID).WillReturnRows(tt.rows)

			got, err := r.Matrix(tt.args.UUID)
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
	validRows := sqlmock.NewRows([]string{
		"uuid", "code", "name", "description", "course_uuid",
		"created_at", "updated_at"}).
		AddRow(
			matrix.UUID, matrix.Code, matrix.Name, matrix.Description, matrix.CourseUUID,
			matrix.CreatedAt, matrix.UpdatedAt).
		AddRow(uuid.MustParse("e74868b2-72d4-4591-a90d-122a9ac2d945"),
			matrix.Code, matrix.Name, matrix.Description, matrix.CourseUUID,
			matrix.CreatedAt, matrix.UpdatedAt)

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
			rows:    utils.EmptyRows,
			wantLen: 0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newMatrixTestDB()
			r, err := NewMatrixRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected when creating the MatrixRepository", err)
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

func TestRepository_CreateMatrix(t *testing.T) {
	validRows := sqlmock.NewRows([]string{
		"uuid", "code", "name", "description", "course_uuid",
		"created_at", "updated_at"}).
		AddRow(
			matrix.UUID, matrix.Code, matrix.Name, matrix.Description, matrix.CourseUUID,
			matrix.CreatedAt, matrix.UpdatedAt)

	type args struct {
		s *domain.Matrix
	}

	tests := []struct {
		name    string
		rows    *sqlmock.Rows
		args    args
		want    domain.Matrix
		wantErr bool
	}{
		{
			name:    "create matrix",
			rows:    validRows,
			args:    args{s: &matrix},
			want:    matrix,
			wantErr: false,
		},
		{
			name:    "empty fields",
			rows:    utils.EmptyRows,
			args:    args{s: &matrix},
			want:    domain.Matrix{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newMatrixTestDB()
			r, err := NewMatrixRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected when creating the repository", err)
			}
			prep, ok := stmts[createMatrix]
			if !ok {
				t.Fatalf("prepared statement %s not found", createMatrix)
			}

			prep.ExpectQuery().WillReturnRows(tt.rows)

			if err := r.CreateMatrix(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("CreateMatrix() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && !reflect.DeepEqual(*tt.args.s, tt.want) {
				t.Errorf("CreateMatrix() got = %v, want %v", *tt.args.s, tt.want)
			}
		})
	}
}

func TestRepository_UpdateMatrix(t *testing.T) {
	validRows := sqlmock.NewRows([]string{
		"uuid", "code", "name", "description", "course_uuid",
		"created_at", "updated_at"}).
		AddRow(
			matrix.UUID, matrix.Code, matrix.Name, matrix.Description, matrix.CourseUUID,
			matrix.CreatedAt, matrix.UpdatedAt)

	type args struct {
		m *domain.Matrix
	}
	tests := []struct {
		name    string
		args    args
		rows    *sqlmock.Rows
		want    domain.Matrix
		wantErr bool
	}{
		{
			name:    "update matrix",
			args:    args{m: &matrix},
			rows:    validRows,
			want:    matrix,
			wantErr: false,
		},
		{
			name:    "empty matrix",
			args:    args{m: &domain.Matrix{}},
			rows:    utils.EmptyRows,
			want:    domain.Matrix{},
			wantErr: true,
		},
	}
	for _, testCase := range tests {
		tt := testCase
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newMatrixTestDB()
			r, err := NewMatrixRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected creating the MatrixRepository", err)
			}
			prep, ok := stmts[updateMatrix]
			if !ok {
				t.Fatalf("prepared statement %s not found", updateMatrix)
			}

			prep.ExpectQuery().WillReturnRows(tt.rows)

			m, err := r.UpdateMatrix(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateMatrix() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && !reflect.DeepEqual(m, tt.want) {
				t.Errorf("UpdateMatrix() got = %v, want %v", *tt.args.m, tt.want)
			}
		})
	}
}
