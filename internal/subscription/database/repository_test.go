package database

import (
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/subscription/domain"
	mtests "github.com/sumelms/microservice-course/tests/database"
)

var now = time.Now()

func newTestSubscription() domain.Subscription {
	return domain.Subscription{
		ID:         1,
		UUID:       uuid.MustParse("dd7c915b-849a-4ba4-bc09-aeecd95c40cc"),
		UserID:     uuid.MustParse("ef2bc01e-be93-4a1f-9e96-c78d3d432088"),
		CourseID:   uuid.MustParse("e8276e31-9a87-4cf1-a16c-080f9c5790d1"),
		MatrixID:   uuid.MustParse("0ac0fe6f-4f34-468d-84f9-9e4fc56b0135"),
		ValidUntil: &now,
		CreatedAt:  now,
		UpdatedAt:  now,
		DeletedAt:  nil,
	}
}

func TestRepository_Subscription(t *testing.T) {
	db, mock := mtests.NewDBMock()

	s := newTestSubscription()
	rows := mock.NewRows([]string{"id", "uuid", "user_id", "course_id", "matrix_id",
		"valid_until", "created_at", "updated_at", "deleted_at"}).
		AddRow(s.ID, s.UUID, s.UserID, s.CourseID, s.MatrixID,
			s.ValidUntil, s.CreatedAt, s.UpdatedAt, s.DeletedAt)

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
		want    domain.Subscription
		wantErr bool
	}{
		{
			name:    "get subscription",
			fields:  fields{DB: db},
			args:    args{id: s.UUID},
			rows:    rows,
			want:    s,
			wantErr: false,
		},
		{
			name:    "course not found error",
			fields:  fields{DB: db},
			args:    args{id: uuid.MustParse("8281f61e-956e-4f64-ac0e-860c444c5f86")},
			rows:    rows,
			want:    domain.Subscription{},
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

			query := "SELECT \\* FROM subscriptions WHERE deleted_at IS NULL AND uuid = \\$1"
			mock.ExpectQuery(query).WithArgs(tt.args.id).WillReturnRows(tt.rows)

			got, err := r.Subscription(tt.args.id)
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

func TestRepository_Subscriptions(t *testing.T) {
	db, mock := mtests.NewDBMock()
	s := newTestSubscription()

	rows := mock.NewRows([]string{"id", "uuid", "user_id", "course_id", "matrix_id", "valid_until",
		"created_at", "updated_at", "deleted_at"}).
		AddRow(s.ID, s.UUID, s.UserID, s.CourseID, s.MatrixID, s.ValidUntil,
			s.CreatedAt, s.UpdatedAt, s.DeletedAt).
		AddRow(2, uuid.MustParse("7aec21ad-2fa8-4ddd-b5af-073144031ecc"), s.UserID,
			s.CourseID, s.MatrixID, s.ValidUntil, s.CreatedAt, s.UpdatedAt, s.DeletedAt)

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
			name:    "get all subscriptions",
			fields:  fields{DB: db},
			rows:    rows,
			wantLen: 2,
			wantErr: false,
		},
		{
			name:    "get no subscriptions",
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

			query := "SELECT \\* FROM subscriptions WHERE deleted_at IS NULL"
			mock.ExpectQuery(query).WillReturnRows(rows)

			got, err := r.Subscriptions()
			if (err != nil) != tt.wantErr {
				t.Errorf("Subscriptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("Subscriptions() got = %v, want %v", got, tt.wantLen)
			}
		})
	}
}

func TestRepository_CreateSubscription(t *testing.T) {
	db, mock := mtests.NewDBMock()

	s := newTestSubscription()
	rows := mock.NewRows([]string{"id", "uuid", "user_id", "course_id", "matrix_id",
		"valid_until", "created_at", "updated_at", "deleted_at"}).
		AddRow(s.ID, s.UUID, s.UserID, s.CourseID, s.MatrixID,
			s.ValidUntil, s.CreatedAt, s.UpdatedAt, s.DeletedAt)

	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		c *domain.Subscription
	}

	tests := []struct {
		name    string
		fields  fields
		rows    *sqlmock.Rows
		args    args
		wantErr bool
	}{
		{
			name:    "create subscription",
			fields:  fields{DB: db},
			rows:    rows,
			args:    args{c: &s},
			wantErr: false,
		},
		{
			name:    "empty fields",
			fields:  fields{DB: db},
			rows:    nil,
			args:    args{c: &s},
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

			query := "INSERT INTO subscriptions \\(course_id, matrix_id, user_id, valid_until\\) VALUES \\(\\$1, \\$2, \\$3, \\$4\\) RETURNING \\*"
			mock.ExpectQuery(query).WithArgs(s.CourseID, s.MatrixID, s.UserID, s.ValidUntil).WillReturnRows(rows)

			if err := r.CreateSubscription(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CreateSubscription() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_UpdateSubscription(t *testing.T) {
	db, mock := mtests.NewDBMock()

	s := newTestSubscription()
	rows := mock.NewRows([]string{"id", "uuid", "user_id", "course_id", "matrix_id",
		"valid_until", "created_at", "updated_at", "deleted_at"}).
		AddRow(s.ID, s.UUID, s.UserID, s.CourseID, s.MatrixID,
			s.ValidUntil, s.CreatedAt, s.UpdatedAt, s.DeletedAt)

	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		c *domain.Subscription
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
			args:    args{c: &s},
			rows:    rows,
			wantErr: false,
		},
		{
			name:    "empty course",
			fields:  fields{DB: db},
			args:    args{c: &domain.Subscription{}},
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

			// nolint: lll
			query := `UPDATE subscriptions SET user_id = \\$1, course_id = \\$2, matrix_id = \\$3, valid_until = \\$4 WHERE uuid = \\$5 RETURNING \\*`
			mock.ExpectQuery(query).WithArgs(s.UserID, s.CourseID, s.MatrixID, s.ValidUntil, s.UUID).WillReturnRows(rows)

			if err := r.UpdateSubscription(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UpdateSubscription() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
