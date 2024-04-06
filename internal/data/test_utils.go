package data

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

// add all the seeds to this single query
const seeds = `
INSERT INTO users (name, age, birthdate, gender, location, phone) VALUES ('test_user', 22, DATE '2022-10-06', 'male', '78.123,11.123', '9999999999');
`

// cleanup test database
func cleanupDB(db *sqlx.DB) {
	for _, t := range tables {
		_, _ = db.Exec(fmt.Sprintf("DELETE FROM %s", t))
	}

	db.Close()
}

func parseDate(format, value string) time.Time {
	res, _ := time.ParseInLocation(format, value, &time.Location{})
	return res
}
