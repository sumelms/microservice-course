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
	subject = domain.Subject{
		UUID:        utils.SubjectUUID,
		Code:        "Code",
		Name:        "Name",
		Objective:   "Objective",
		Credit:      10,
		Workload:    20,
		PublishedAt: &utils.Now,
		CreatedAt:   utils.Now,
		UpdatedAt:   utils.Now,
	}
)

func newSubjectTestDB() (*sqlx.DB, sqlmock.Sqlmock, map[string]*sqlmock.ExpectedPrepare) {
	return utils.NewTestDB(queriesSubject())
}

func TestRepository_Subject(t *testing.T) {
	validRows := sqlmock.NewRows([]string{
		"uuid", "code", "name", "objective", "credit", "workload",
		"created_at", "updated_at", "published_at"}).
		AddRow(
			subject.UUID, subject.Code, subject.Name, subject.Objective, subject.Credit, subject.Workload,
			subject.CreatedAt, subject.UpdatedAt, subject.PublishedAt)

	type args struct {
		UUID uuid.UUID
	}

	tests := []struct {
		name    string
		args    args
		rows    *sqlmock.Rows
		want    domain.Subject
		wantErr bool
	}{
		{
			name:    "get subject",
			args:    args{UUID: utils.SubjectUUID},
			rows:    validRows,
			want:    subject,
			wantErr: false,
		},
		{
			name:    "subject not found error",
			args:    args{UUID: uuid.MustParse("8281f61e-956e-4f64-ac0e-860c444c5f86")},
			rows:    utils.EmptyRows,
			want:    domain.Subject{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newSubjectTestDB()
			r, err := NewSubjectRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected when creating the SubjectRepository", err)
			}
			prep, ok := stmts[getSubject]
			if !ok {
				t.Fatalf("prepared statement %s not found", getSubject)
			}

			prep.ExpectQuery().WithArgs(utils.SubjectUUID).WillReturnRows(tt.rows)

			got, err := r.Subject(tt.args.UUID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_Subjects(t *testing.T) {
	validRows := sqlmock.NewRows([]string{
		"uuid", "code", "name", "objective", "credit", "workload",
		"created_at", "updated_at", "published_at"}).
		AddRow(
			subject.UUID, subject.Code, subject.Name, subject.Objective, subject.Credit, subject.Workload,
			subject.CreatedAt, subject.UpdatedAt, subject.PublishedAt).
		AddRow(uuid.MustParse("e74868b2-72d4-4591-a90d-122a9ac2d945"),
			subject.Code, subject.Name, subject.Objective, subject.Credit, subject.Workload,
			subject.CreatedAt, subject.UpdatedAt, subject.PublishedAt)

	tests := []struct {
		name    string
		rows    *sqlmock.Rows
		wantLen int
		wantErr bool
	}{
		{
			name:    "get all subjects",
			rows:    validRows,
			wantLen: 2,
			wantErr: false,
		},
		{
			name:    "get no subjects",
			rows:    utils.EmptyRows,
			wantLen: 0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newSubjectTestDB()
			r, err := NewSubjectRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected when creating the SubjectRepository", err)
			}
			prep, ok := stmts[listSubjects]
			if !ok {
				t.Fatalf("prepared statement %s not found", listSubjects)
			}

			prep.ExpectQuery().WillReturnRows(tt.rows)

			got, err := r.Subjects()
			if (err != nil) != tt.wantErr {
				t.Errorf("Subjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("Subjects() got = %v, want %v", got, tt.wantLen)
			}
		})
	}
}

func TestRepository_CreateSubject(t *testing.T) {
	validRows := sqlmock.NewRows([]string{
		"uuid", "code", "name", "objective", "credit", "workload",
		"created_at", "updated_at", "published_at"}).
		AddRow(
			subject.UUID, subject.Code, subject.Name, subject.Objective, subject.Credit, subject.Workload,
			subject.CreatedAt, subject.UpdatedAt, subject.PublishedAt)

	type args struct {
		subj *domain.Subject
	}

	tests := []struct {
		name    string
		rows    *sqlmock.Rows
		args    args
		want    domain.Subject
		wantErr bool
	}{
		{
			name:    "create subject",
			rows:    validRows,
			args:    args{subj: &subject},
			want:    subject,
			wantErr: false,
		},
		{
			name:    "empty fields",
			rows:    utils.EmptyRows,
			args:    args{subj: &subject},
			want:    domain.Subject{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newSubjectTestDB()
			r, err := NewSubjectRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected when creating the repository", err)
			}
			prep, ok := stmts[createSubject]
			if !ok {
				t.Fatalf("prepared statement %s not found", createSubject)
			}

			prep.ExpectQuery().WillReturnRows(tt.rows)

			if err := r.CreateSubject(tt.args.subj); (err != nil) != tt.wantErr {
				t.Errorf("CreateSubject() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && !reflect.DeepEqual(*tt.args.subj, tt.want) {
				t.Errorf("CreateSubject() got = %v, want %v", *tt.args.subj, tt.want)
			}
		})
	}
}

func TestRepository_UpdateSubject(t *testing.T) {
	validRows := sqlmock.NewRows([]string{
		"uuid", "code", "name", "objective", "credit", "workload",
		"published_at", "created_at", "updated_at"}).
		AddRow(
			subject.UUID, subject.Code, subject.Name, subject.Objective, subject.Credit, subject.Workload,
			subject.PublishedAt, subject.CreatedAt, subject.UpdatedAt)

	type args struct {
		subj *domain.Subject
	}
	tests := []struct {
		name    string
		args    args
		rows    *sqlmock.Rows
		want    domain.Subject
		wantErr bool
	}{
		{
			name:    "update subject",
			args:    args{subj: &subject},
			rows:    validRows,
			want:    subject,
			wantErr: false,
		},
		{
			name:    "empty subject",
			args:    args{subj: &domain.Subject{}},
			rows:    utils.EmptyRows,
			want:    domain.Subject{},
			wantErr: true,
		},
	}
	for _, testCase := range tests {
		tt := testCase
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newSubjectTestDB()
			r, err := NewSubjectRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected creating the SubjectRepository", err)
			}
			prep, ok := stmts[updateSubject]
			if !ok {
				t.Fatalf("prepared statement %s not found", updateSubject)
			}
			prep.ExpectQuery().WillReturnRows(tt.rows)

			if err := r.UpdateSubject(tt.args.subj); (err != nil) != tt.wantErr {
				t.Errorf("UpdateSubject() \nerror = %v, \nwantErr = %v", err, tt.wantErr)
			}

			if !tt.wantErr && !reflect.DeepEqual(*tt.args.subj, tt.want) {
				t.Errorf("UpdateSubject() \ngot = %v, \nwant = %v", *tt.args.subj, tt.want)
			}
		})
	}
}

func TestRepository_DeleteSubject(t *testing.T) {
	validRows := sqlmock.NewRows([]string{"uuid", "deleted_at"}).
		AddRow(subject.UUID, utils.Now)

	deletedSubject := domain.DeletedSubject{
		UUID:      utils.SubjectUUID,
		DeletedAt: utils.Now,
	}

	type args struct {
		subj *domain.DeletedSubject
	}
	tests := []struct {
		name    string
		args    args
		rows    *sqlmock.Rows
		want    domain.DeletedSubject
		wantErr bool
	}{
		{
			name:    "delete subject",
			args:    args{subj: &deletedSubject},
			rows:    validRows,
			want:    deletedSubject,
			wantErr: false,
		},
		{
			name:    "empty subject",
			args:    args{subj: &domain.DeletedSubject{}},
			rows:    utils.EmptyRows,
			want:    domain.DeletedSubject{},
			wantErr: true,
		},
	}
	for _, testCase := range tests {
		tt := testCase
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newSubjectTestDB()
			r, err := NewSubjectRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected creating the SubjectRepository", err)
			}

			prep, ok := stmts[deleteSubject]
			if !ok {
				t.Fatalf("prepared statement %s not found", deleteSubject)
			}
			prep.ExpectQuery().WillReturnRows(tt.rows)

			if err := r.DeleteSubject(tt.args.subj); (err != nil) != tt.wantErr {
				t.Errorf("DeleteSubject() \nerror = %v, \nwantErr = %v", err, tt.wantErr)
			}

			if !tt.wantErr && !reflect.DeepEqual(*tt.args.subj, tt.want) {
				t.Errorf("DeleteSubject() \ngot = %v, \nwant = %v", *tt.args.subj, tt.want)
			}
		})
	}
}
