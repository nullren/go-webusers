package memstore

import (
	"errors"
	"sync"
	"sync/atomic"

	"github.com/nullren/go-webusers/pkg/user"
)

type LocalStorage struct {
	mux         sync.Mutex
	users       uint64
	idxUsername map[string]*user.User
	idxID       map[user.ID]*user.User
}

func NewStore() *LocalStorage {
	return &LocalStorage{
		users:       0,
		idxUsername: make(map[string]*user.User),
		idxID:       make(map[user.ID]*user.User),
	}
}

func (l LocalStorage) Create(newUser user.User) (*user.User, error) {
	if _, ok := l.idxUsername[newUser.Username]; ok {
		return nil, ErrUserExists
	}

	l.mux.Lock()
	defer l.mux.Unlock()

	if _, ok := l.idxUsername[newUser.Username]; ok {
		return nil, ErrUserExists
	}

	id := atomic.AddUint64(&l.users, 1)
	newUser.ID = user.ID(id)

	l.idxUsername[newUser.Username] = &newUser
	l.idxID[newUser.ID] = &newUser

	return &newUser, nil
}

func (l LocalStorage) Auth(user user.User, compare user.ComparePasswordHash) (*user.User, error) {
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

var _ user.Store = (*LocalStorage)(nil)
