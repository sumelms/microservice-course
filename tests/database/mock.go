package database

import (
	"fmt"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

type mockMatcher struct{}

func (m mockMatcher) Match(expectedSQL, actualSQL string) error {
	if expectedSQL != actualSQL {
		return fmt.Errorf("actual sql %s is not as expected %s", actualSQL, expectedSQL)
	}
	return nil
}

func NewDBMock() (*sqlx.DB, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(mockMatcher{}))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db := sqlx.NewDb(mockDB, "sqlmock")
	return db, mock
}
