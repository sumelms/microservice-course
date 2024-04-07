package database

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/subscription/domain"
	utils "github.com/sumelms/microservice-course/tests"
)

var (
	subscription = domain.Subscription{
		UUID:       utils.SubscriptionUUID,
		UserUUID:   utils.UserUUID,
		CourseUUID: &utils.CourseUUID,
		MatrixUUID: &utils.MatrixUUID,
		Role:       utils.Role,
		ExpiresAt:  &utils.Now,
		CreatedAt:  utils.Now,
		UpdatedAt:  utils.Now,
	}
)

func newSubscriptionTestDB() (*sqlx.DB, sqlmock.Sqlmock, map[string]*sqlmock.ExpectedPrepare) {
	return utils.NewTestDB(queriesSubscription())
}

func TestRepository_Subscription(t *testing.T) {
	validRows := sqlmock.NewRows([]string{"uuid", "user_uuid", "course_uuid", "matrix_uuid", "role",
		"expires_at", "created_at", "updated_at"}).
		AddRow(subscription.UUID, subscription.UserUUID, subscription.CourseUUID, subscription.MatrixUUID, subscription.Role,
			subscription.ExpiresAt, subscription.CreatedAt, subscription.UpdatedAt)

	type args struct {
		UUID uuid.UUID
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
			args:    args{UUID: utils.SubscriptionUUID},
			rows:    validRows,
			want:    subscription,
			wantErr: false,
		},
		{
			name:    "course not found error",
			args:    args{UUID: uuid.MustParse("00000000-0000-0000-0000-000000000000")},
			rows:    utils.EmptyRows,
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

			prep.ExpectQuery().WithArgs(utils.SubscriptionUUID).WillReturnRows(tt.rows)

			got, err := r.Subscription(tt.args.UUID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subscription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subscription() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_Subscriptions(t *testing.T) {
	validRows := sqlmock.NewRows([]string{"uuid", "user_uuid", "course_uuid", "matrix_uuid", "role",
		"expires_at", "created_at", "updated_at"}).
		AddRow(subscription.UUID, subscription.UserUUID, subscription.CourseUUID, subscription.MatrixUUID, subscription.Role,
			subscription.ExpiresAt, subscription.CreatedAt, subscription.UpdatedAt).
		AddRow(uuid.MustParse("7aec21ad-2fa8-4ddd-b5af-073144031ecc"), subscription.UserUUID, subscription.CourseUUID, subscription.MatrixUUID, subscription.Role,
			subscription.ExpiresAt, subscription.CreatedAt, subscription.UpdatedAt)

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
			rows:    utils.EmptyRows,
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

func TestRepository_CreateSubscriptionWithoutMatrix(t *testing.T) {
	validRows := sqlmock.NewRows([]string{
		"uuid", "user_uuid", "role",
		"expires_at", "created_at", "updated_at"}).
		AddRow(
			subscription.UUID, subscription.UserUUID, subscription.Role,
			subscription.ExpiresAt, subscription.CreatedAt, subscription.UpdatedAt)

	type args struct {
		s *domain.Subscription
	}

	tests := []struct {
		name    string
		rows    *sqlmock.Rows
		args    args
		want    domain.Subscription
		wantErr bool
	}{
		{
			name:    "create subscription",
			rows:    validRows,
			args:    args{s: &subscription},
			want:    subscription,
			wantErr: false,
		},
		{
			name:    "empty fields",
			rows:    utils.EmptyRows,
			args:    args{s: &subscription},
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
				t.Fatalf("an error '%s' was not expected when creating the repository", err)
			}
			prep, ok := stmts[createSubscriptionWithoutMatrix]
			if !ok {
				t.Fatalf("prepared statement %s not found", createSubscriptionWithoutMatrix)
			}

			prep.ExpectQuery().WillReturnRows(tt.rows)

			if err := r.CreateSubscriptionWithoutMatrix(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("CreateSubscription() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && !reflect.DeepEqual(*tt.args.s, tt.want) {
				t.Errorf("CreateSubscription() got = %v, want %v", *tt.args.s, tt.want)
			}
		})
	}
}

func TestRepository_CreateSubscription(t *testing.T) {
	validRows := sqlmock.NewRows([]string{
		"uuid", "user_uuid", "role",
		"expires_at", "created_at", "updated_at"}).
		AddRow(
			subscription.UUID, subscription.UserUUID, subscription.Role,
			subscription.ExpiresAt, subscription.CreatedAt, subscription.UpdatedAt)

	type args struct {
		s *domain.Subscription
	}

	tests := []struct {
		name    string
		rows    *sqlmock.Rows
		args    args
		want    domain.Subscription
		wantErr bool
	}{
		{
			name:    "create subscription",
			rows:    validRows,
			args:    args{s: &subscription},
			want:    subscription,
			wantErr: false,
		},
		{
			name:    "empty fields",
			rows:    utils.EmptyRows,
			args:    args{s: &subscription},
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

			if !tt.wantErr && !reflect.DeepEqual(*tt.args.s, tt.want) {
				t.Errorf("CreateSubscription() got = %v, want %v", *tt.args.s, tt.want)
			}
		})
	}
}

func TestRepository_UpdateSubscription(t *testing.T) {
	validRows := sqlmock.NewRows([]string{
		"uuid", "role",
		"expires_at", "created_at", "updated_at"}).
		AddRow(
			subscription.UUID, subscription.Role,
			subscription.ExpiresAt, subscription.CreatedAt, subscription.UpdatedAt)

	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		s *domain.Subscription
	}
	tests := []struct {
		name    string
		fields  fields
		rows    *sqlmock.Rows
		args    args
		want    domain.Subscription
		wantErr bool
	}{
		{
			name:    "update course",
			rows:    validRows,
			args:    args{s: &subscription},
			want:    subscription,
			wantErr: false,
		},
		{
			name:    "empty course",
			rows:    utils.EmptyRows,
			args:    args{s: &domain.Subscription{}},
			want:    domain.Subscription{},
			wantErr: true,
		},
	}
	for _, testCase := range tests {
		tt := testCase
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

			if !tt.wantErr && !reflect.DeepEqual(*tt.args.s, tt.want) {
				t.Errorf("CreateSubscription() got = %v, want %v", *tt.args.s, tt.want)
			}
		})
	}
}
