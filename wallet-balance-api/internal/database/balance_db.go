package database

import (
	"database/sql"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-event-driven-architecture/wallet-balance-api/internal/entity"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(DB *sql.DB) *BalanceDB {
	return &BalanceDB{DB: DB}
}

func (db *BalanceDB) FindByAccountID(accountID string) (*entity.Balance, error) {
	var balance entity.Balance

	stmt, err := db.DB.Prepare("SELECT id, account_id, balance, created_at, updated_at FROM balances WHERE account_id = ? ORDER BY id DESC")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(accountID)

	err = row.Scan(&balance.ID, &balance.AccountID, &balance.Balance, &balance.CreatedAt, &balance.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &balance, nil
}

func (db *BalanceDB) Save(balance *entity.Balance) error {
	stmt, err := db.DB.Prepare("INSERT INTO balances (id, account_id, balance, created_at, updated_at) VALUES (?, ?, ?, NOW(), NOW())")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(balance.ID, balance.AccountID, balance.Balance)

	if err != nil {
		return err
	}

	return nil
}
