package webusers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/nullren/go-webusers/pkg/session"
	"github.com/nullren/go-webusers/pkg/user"
)

type Handlers struct {
	Users    *user.Controller
	Sessions *session.Controller
}

func (h Handlers) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username, password, err := readAuth(r)
		if err != nil {
			fail(w, err)
			return
		}
		newUser, err := h.Users.SignUp(username, password)
		if err != nil {
			fail(w, err)
			return
		}
		log.Printf("created user: %s", newUser)
		h.Sessions.Write(w, newUser)
		http.Redirect(w, r, "/settings", 301)
		return
	}
	render(w, "signup.html")
}

func (h Handlers) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username, password, err := readAuth(r)
		if err != nil {
			fail(w, err)
			return
		}
		loggedIn, err := h.Users.Login(username, password)
		if err != nil {
			fail(w, err)
			return
		}
		log.Printf("logged in user: %s", loggedIn)
		h.Sessions.Write(w, loggedIn)
		http.Redirect(w, r, "/settings", 301)
		return
	}
	render(w, "login.html")
}

func (h Handlers) Settings(w http.ResponseWriter, r *http.Request) {
	render(w, "settings.html")
}

func fail(w http.ResponseWriter, err error) {
	log.Printf("failure: %s\n", err)
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(fmt.Sprintf("500 - %s", err)))
}

func readAuth(r *http.Request) (string, string, error) {
	if err := r.ParseForm(); err != nil {
		return "", "", err
	}
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	return username, password, nil
}

func render(w http.ResponseWriter, fileName string) {
	lp := filepath.Join("web", "template", "layout.html")
	fp := filepath.Join("web", "template", fileName)
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		log.Println(path)
		log.Printf("error parsing template: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("500 - Something bad happened!"))
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("500 - Something bad happened!"))
	}
}
