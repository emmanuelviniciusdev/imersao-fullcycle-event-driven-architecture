package migration

import (
	"database/sql"
)

func RunMigrationAccount(db *sql.DB) {
	db.Exec("CREATE TABLE IF NOT EXISTS accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")

	db.Exec("TRUNCATE TABLE accounts")

	db.Exec("INSERT INTO accounts (id, client_id, balance, created_at) VALUES ('2312fd16-683f-47dc-a4c7-e31985727d73', '46ce90d3-a3e6-42a3-a10d-4a3e7e0e2866', 100, NOW())")
	db.Exec("INSERT INTO accounts (id, client_id, balance, created_at) VALUES ('dc6e87db-6e52-48d6-8a42-c9879932429e', '6c7b9331-cfea-439c-b752-e4ac10707ab1', 100, NOW())")
}
