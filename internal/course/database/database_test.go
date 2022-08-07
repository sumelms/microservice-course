package database

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
	now              = time.Now()
	courseUUID       = uuid.MustParse("e8276e31-9a87-4cf1-a16c-080f9c5790d1")
	subscriptionUUID = uuid.MustParse("dd7c915b-849a-4ba4-bc09-aeecd95c40cc")
	userUUID         = uuid.MustParse("ef2bc01e-be93-4a1f-9e96-c78d3d432088")
	matrixUUID       = uuid.MustParse("0ac0fe6f-4f34-468d-84f9-9e4fc56b0135")
	emptyRows        = sqlmock.NewRows([]string{})
)

func newTestDB(queries map[string]string) (*sqlx.DB, sqlmock.Sqlmock, map[string]*sqlmock.ExpectedPrepare) {
	db, mock := database.NewDBMock()

	sqlStatements := make(map[string]*sqlmock.ExpectedPrepare)
	for queryName, query := range queries {
		stmt := mock.ExpectPrepare(fmt.Sprintf("^%s$", regexp.QuoteMeta(string(query))))
		sqlStatements[queryName] = stmt
	}

	mock.MatchExpectationsInOrder(false)
	return db, mock, sqlStatements
}
