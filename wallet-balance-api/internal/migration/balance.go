package migration

import (
	"database/sql"
)

func RunMigrationBalance(db *sql.DB) {
	db.Exec("CREATE TABLE IF NOT EXISTS balances (id int not null unique auto_increment, account_id varchar(255), balance double, created_at date, updated_at date)")

	db.Exec("TRUNCATE TABLE balances")
}
