package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type ID uint64

type User struct {
	ID           ID
	Username     string
	PasswordHash []byte
}

func (u User) String() string {
	return fmt.Sprintf("<user ID:%d>", u.ID)
}

type ComparePasswordHash func([]byte) error

type Store interface {
	Create(User) (*User, error)
	Auth(User, ComparePasswordHash) (*User, error)
}

type Controller struct {
	dao Store
}

func New(dao Store) *Controller {
	return &Controller{dao}
}

func (c *Controller) SignUp(username, password string) (*User, error) {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	return c.dao.Create(User{
		ID:           0,
		Username:     username,
		PasswordHash: pwdHash,
	})
}

func (c *Controller) Login(username, password string) (*User, error) {
	compare := func(h []byte) error {
		return bcrypt.CompareHashAndPassword(h, []byte(password))
	}

	return c.dao.Auth(User{
		ID:           0,
		Username:     username,
		PasswordHash: nil,
	}, compare)
}
