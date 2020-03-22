package webusers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/nullren/go-webusers/pkg/user"
)

type Handlers struct {
	Users *user.Controller
}

func (h Handlers) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username, password, err := readAuth(r)
		if err != nil {
			fail(w, err)
			return
		}
		log.Printf("read user: %q, pass: %q\n", username, password)
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
		log.Printf("read user: %q, pass: %q\n", username, password)
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
