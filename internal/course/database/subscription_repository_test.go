package database

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/course/domain"
)

var (
	subscription = domain.Subscription{
		ID:         1,
		UUID:       subscriptionUUID,
		UserID:     userUUID,
		CourseID:   courseUUID,
		MatrixID:   &matrixUUID,
		ValidUntil: &now,
		CreatedAt:  now,
		UpdatedAt:  now,
		DeletedAt:  nil,
	}
)

func newSubscriptionTestDB() (*sqlx.DB, sqlmock.Sqlmock, map[string]*sqlmock.ExpectedPrepare) {
	return newTestDB(subscriptionQueries())
}

func TestRepository_Subscription(t *testing.T) {
	validRows := sqlmock.NewRows([]string{"id", "uuid", "user_id", "course_id", "matrix_id",
		"valid_until", "created_at", "updated_at", "deleted_at"}).
		AddRow(subscription.ID, subscription.UUID, subscription.UserID, subscription.CourseID, subscription.MatrixID,
			subscription.ValidUntil, subscription.CreatedAt, subscription.UpdatedAt, subscription.DeletedAt)

	type args struct {
		id uuid.UUID
	}

	tests := []struct {
		name    string
		args    args
		rows    *sqlmock.Rows
		want    domain.Subscription
		wantErr bool
	}{
		{
			name:    "get subscription",
			args:    args{id: subscriptionUUID},
			rows:    validRows,
			want:    subscription,
			wantErr: false,
		},
		{
			name:    "course not found error",
			args:    args{id: uuid.MustParse("8281f61e-956e-4f64-ac0e-860c444c5f86")},
			rows:    emptyRows,
			want:    domain.Subscription{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newSubscriptionTestDB()
			r, err := NewSubscriptionRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected creating the repository", err)
			}
			prep, ok := stmts[getSubscription]
			if !ok {
				t.Fatalf("prepared statement %s not found", getSubscription)
			}

			prep.ExpectQuery().WithArgs(subscriptionUUID).WillReturnRows(tt.rows)

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
	validRows := sqlmock.NewRows([]string{"id", "uuid", "user_id", "course_id", "matrix_id", "valid_until",
		"created_at", "updated_at", "deleted_at"}).
		AddRow(subscription.ID, subscription.UUID, subscription.UserID, subscription.CourseID, subscription.MatrixID,
			subscription.ValidUntil, subscription.CreatedAt, subscription.UpdatedAt, subscription.DeletedAt).
		AddRow(2, uuid.MustParse("7aec21ad-2fa8-4ddd-b5af-073144031ecc"), subscription.UserID,
			subscription.CourseID, subscription.MatrixID, subscription.ValidUntil, subscription.CreatedAt,
			subscription.UpdatedAt, subscription.DeletedAt)

	tests := []struct {
		name    string
		rows    *sqlmock.Rows
		wantLen int
		wantErr bool
	}{
		{
			name:    "get all subscriptions",
			rows:    validRows,
			wantLen: 2,
			wantErr: false,
		},
		{
			name:    "get no subscriptions",
			rows:    emptyRows,
			wantLen: 0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newSubscriptionTestDB()
			r, err := NewSubscriptionRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected creating the repository", err)
			}
			prep, ok := stmts[listSubscription]
			if !ok {
				t.Fatalf("prepared statement %s not found", listSubscription)
			}

			prep.ExpectQuery().WillReturnRows(tt.rows)

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
	validRows := sqlmock.NewRows([]string{"id", "uuid", "user_id", "course_id", "matrix_id",
		"valid_until", "created_at", "updated_at", "deleted_at"}).
		AddRow(subscription.ID, subscription.UUID, subscription.UserID, subscription.CourseID, subscription.MatrixID,
			subscription.ValidUntil, subscription.CreatedAt, subscription.UpdatedAt, subscription.DeletedAt)

	type args struct {
		s *domain.Subscription
	}

	tests := []struct {
		name    string
		rows    *sqlmock.Rows
		args    args
		wantErr bool
	}{
		{
			name:    "create subscription",
			rows:    validRows,
			args:    args{s: &subscription},
			wantErr: false,
		},
		{
			name:    "empty fields",
			rows:    emptyRows,
			args:    args{s: &subscription},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newSubscriptionTestDB()
			r, err := NewSubscriptionRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected when creating the repository", err)
			}
			prep, ok := stmts[createSubscription]
			if !ok {
				t.Fatalf("prepared statement %s not found", createSubscription)
			}

			prep.ExpectQuery().WillReturnRows(tt.rows)

			if err := r.CreateSubscription(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("CreateSubscription() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_UpdateSubscription(t *testing.T) {
	validRows := sqlmock.NewRows([]string{"id", "uuid", "user_id", "course_id", "matrix_id",
		"valid_until", "created_at", "updated_at", "deleted_at"}).
		AddRow(subscription.ID, subscription.UUID, subscription.UserID, subscription.CourseID, subscription.MatrixID,
			subscription.ValidUntil, subscription.CreatedAt, subscription.UpdatedAt, subscription.DeletedAt)

	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		s *domain.Subscription
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
			args:    args{s: &subscription},
			rows:    validRows,
			wantErr: false,
		},
		{
			name:    "empty course",
			args:    args{s: &domain.Subscription{}},
			rows:    emptyRows,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newSubscriptionTestDB()
			r, err := NewSubscriptionRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected when creating the repository", err)
			}
			prep, ok := stmts[updateSubscription]
			if !ok {
				t.Fatalf("prepared statement %s not found", updateSubscription)
			}

			prep.ExpectQuery().WillReturnRows(tt.rows)

			if err := r.UpdateSubscription(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("UpdateSubscription() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
