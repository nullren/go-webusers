package user

import (
	"errors"
	"sync/atomic"
)

type LocalStorage struct {
	users       uint64
	idxUsername map[string]*User
	idxID       map[ID]*User
}

func (l LocalStorage) Create(user User) (*User, error) {
	if _, ok := l.idxUsername[user.Username]; ok {
		return nil, ErrUserExists
	}
	id := atomic.AddUint64(&l.users, 1)
	user.ID = ID(id)

	l.idxUsername[user.Username] = &user
	l.idxID[user.ID] = &user

	return &user, nil
}

func (l LocalStorage) Auth(user User, compare ComparePasswordHash) (*User, error) {
	found, ok := l.idxUsername[user.Username]
	if !ok {
		return nil, ErrUserNotFound
	}

	err := compare(found.PasswordHash)
	if err != nil {
		return nil, ErrPasswordIncorrect
	}

	return found, nil
}

var (
	ErrUserExists        = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrPasswordIncorrect = errors.New("user password incorrect")
)

var _ Store = (*LocalStorage)(nil)
