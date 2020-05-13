package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20200512212639, Down20200512212639)
}

func Up20200512212639(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func Down20200512212639(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
