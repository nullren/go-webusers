package webusers

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	render(w, "signup.html")
}

func Login(w http.ResponseWriter, r *http.Request) {
	render(w, "login.html")
}

func Settings(w http.ResponseWriter, r *http.Request) {
	render(w, "settings.html")
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

