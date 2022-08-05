package object

import (
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type (
	AccountID    = int64
	PasswordHash = string

	// Account account
	Account struct {
		// The internal ID of the account
		ID AccountID `json:"-"`

		// The username of the account
		Username string `json:"username,omitempty"`

		// The username of the account
		PasswordHash string `json:"-" db:"password_hash"`

		// The account's display name
		DisplayName *string `json:"display_name,omitempty" db:"display_name"`

		// URL to the avatar image
		Avatar *string `json:"avatar,omitempty"`

		// URL to the header image
		Header *string `json:"header,omitempty"`

		// Biography of user
		Note *string `json:"note,omitempty"`

		// The time the account was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)

func NewAccount(username, password string) (Account, error) {
	if len(username) > 10 {
		return Account{}, errors.New("username is too long")
	}

	if len(password) < 5 {
		return Account{}, errors.New("password is too short")
	}

	account := Account{
		Username: username,
		CreateAt: DateTime{Time: time.Now()},
	}

	if err := account.SetPassword(password); err != nil {
		return Account{}, err
	}

	return account, nil
}

// Check if given password is match to account's password
func (a *Account) CheckPassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.PasswordHash), []byte(pass)) == nil
}

// Hash password and set it to account object
func (a *Account) SetPassword(pass string) error {
	passwordHash, err := generatePasswordHash(pass)
	if err != nil {
		return fmt.Errorf("generate error: %w", err)
	}
	a.PasswordHash = passwordHash
	return nil
}

func generatePasswordHash(pass string) (PasswordHash, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hashing password failed: %w", errors.WithStack(err))
	}
	return PasswordHash(hash), nil
}

func (a *Account) SetDisplayName(displayName string) error {
	if strings.Contains(displayName, "@") {
		return errors.New("invalid display name")
	}
	a.DisplayName = &displayName
	return nil
}
