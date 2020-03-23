package memstore

import (
	"sync"
	"time"

	"github.com/nullren/go-webusers/pkg/session"
	"github.com/nullren/go-webusers/pkg/user"
)

type timeUser struct {
	u   *user.User
	exp time.Time
}

type localSession struct {
	mux        sync.Mutex
	idxSession map[string]timeUser
}

func (l localSession) GetSession(key string) (*user.User, error) {
	p, ok := l.idxSession[key]
	if !ok {
		return nil, nil
	}
	if p.exp.Before(time.Now()) {
		go func() {
			l.mux.Lock()
			defer l.mux.Unlock()
			delete(l.idxSession, key)
		}()
		return nil, nil
	}
	return p.u, nil
}

func (l localSession) SetSession(key string, ttl time.Duration, u *user.User) (*user.User, error) {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.idxSession[key] = timeUser{
		u:   u,
		exp: time.Now().Add(ttl),
	}
	return u, nil
}

func NewStore() *localSession {
	return &localSession{
		idxSession: make(map[string]timeUser),
	}
}

var _ session.Store = (*localSession)(nil)
