package models

import (
	"errors"
	"github.com/farmani/sharebuy/pkg/encryption"
	"time"

	"github.com/farmani/sharebuy/internal/models/enums"

	"golang.org/x/crypto/bcrypt"
)

var AnonymousUser = &User{}

type User struct {
	ID        int64            `json:"id"`
	UUID      int64            `json:"uuid"`
	Name      string           `json:"name"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	Password  password         `json:"-"`
	StripeId  string           `json:"stripe_id"`
	Status    enums.UserStatus `json:"status"`
	CreatedAt time.Time        `json:"created_at"`
	Version   int              `json:"-"`
}

func NewUser(email, pass string) (*User, error) {
	p := password{}
	err := p.set(pass)
	if err != nil {
		return &User{}, err
	}

	u := &User{
		Email:    email,
		Password: p,
	}

	u.Username, err = encryption.GenerateRandomString(6)
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) IsAnonymous() bool {
	return u == AnonymousUser
}

type password struct {
	Plaintext *string
	Hash      []byte
}

// Set calculates the bcrypt hash of a plaintext password, and stores both the has and the
// plaintext versions in the password struct.
func (p *password) set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}

	p.Plaintext = &plaintextPassword
	p.Hash = hash
	return nil
}

// Matches checks whether the provided plaintext password matches the hashed password stored in
// the password struct, returning true if it matches and false otherwise.
func (p *password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.Hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
