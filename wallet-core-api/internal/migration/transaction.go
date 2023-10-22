package migration

import (
	"database/sql"
)

func RunMigrationTransaction(db *sql.DB) {
	db.Exec("CREATE TABLE IF NOT EXISTS transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date)")

	db.Exec("TRUNCATE TABLE transactions")
}
