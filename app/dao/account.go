package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

func (r *account) Create(ctx context.Context, account object.Account) error {
	query := `
INSERT INTO account (username, password_hash, display_name, avatar, header, note, create_at)
VALUES (?, ?, ?, ?, ?, ?, ?);`
	if _, err := r.db.ExecContext(ctx, query, account.Username, account.PasswordHash, account.DisplayName, account.Avatar, account.Header, account.Note, account.CreateAt); err != nil {
		return err
	}
	return nil
}

func (r *account) Update(ctx context.Context, account object.Account) error {
	query := "UPDATE account SET display_name = ? WHERE username = ?;"
	if _, err := r.db.ExecContext(ctx, query, account.DisplayName, account.Username); err != nil {
		return err
	}
	return nil
}
