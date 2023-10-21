package migration

import (
	"database/sql"
)

func RunMigrationClient(db *sql.DB) {
	db.Exec("CREATE TABLE IF NOT EXISTS clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")

	db.Exec("TRUNCATE TABLE clients")

	db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES ('46ce90d3-a3e6-42a3-a10d-4a3e7e0e2866', 'Emmanuel', 'emmanuel@icloud.com', NOW())")
	db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES ('6c7b9331-cfea-439c-b752-e4ac10707ab1', 'Vinícius', 'vinicius@icloud.com', NOW())")
}
