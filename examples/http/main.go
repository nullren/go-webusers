package main

import (
	"log"
	"net/http"

	"github.com/nullren/go-webusers/pkg/webusers"
)

func main() {
	http.HandleFunc("/signup", webusers.Signup)
	http.HandleFunc("/login", webusers.Login)
	http.HandleFunc("/settings", webusers.Settings)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

