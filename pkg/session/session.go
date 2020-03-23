package session

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/nullren/go-webusers/pkg/user"
)

type Controller struct {
	dao Store
}

func New(dao Store) *Controller {
	return &Controller{dao}
}

func (c Controller) Write(w http.ResponseWriter, u *user.User) {
	key := uuid.New().String()
	ttl := 30 * time.Minute
	if _, err := c.dao.SetSession(key, ttl, u); err != nil {
		log.Println("failed to store session key", err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   key,
		Expires: time.Now().Add(ttl),
	})
}

func (c Controller) Read(r *http.Request) *user.User {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return nil
	}
	u, err := c.dao.GetSession(cookie.Value)
	if err != nil {
		log.Println("failed to get session key", err)
		return nil
	}
	return u
}

type Store interface {
	GetSession(key string) (*user.User, error)
	SetSession(key string, ttl time.Duration, u *user.User) (*user.User, error)
}
