package tests

import (
	"fmt"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sumelms/microservice-course/tests/database"
)

var (
	Now              = time.Now()
	CourseUUID       = uuid.MustParse("00000000-0000-0000-0000-aaaaaaaaaaaa")
	MatrixUUID       = uuid.MustParse("00000000-0000-0000-0000-bbbbbbbbbbbb")
	SubscriptionUUID = uuid.MustParse("00000000-0000-0000-0000-cccccccccccc")
	SubjectUUID      = uuid.MustParse("00000000-0000-0000-0000-dddddddddddd")
	UserUUID         = uuid.MustParse("00000000-0000-0000-0000-111111111111")
	Role             = "Role"
	EmptyRows        = sqlmock.NewRows([]string{})
)

func NewTestDB(queries map[string]string) (*sqlx.DB, sqlmock.Sqlmock, map[string]*sqlmock.ExpectedPrepare) {
	db, mock := database.NewDBMock()

	sqlStatements := make(map[string]*sqlmock.ExpectedPrepare)
	for queryName, query := range queries {
		stmt := mock.ExpectPrepare(fmt.Sprintf("^%s$", regexp.QuoteMeta(query)))
		sqlStatements[queryName] = stmt
	}

	mock.MatchExpectationsInOrder(false)

	return db, mock, sqlStatements
}
