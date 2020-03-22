package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nullren/go-webusers/pkg/webusers"
)

func main() {
	r := gin.Default()

	r.GET("/signup", gin.WrapF(webusers.Signup))
	r.POST("/signup", gin.WrapF(webusers.Signup))

	r.GET("/login", gin.WrapF(webusers.Login))
	r.POST("/login", gin.WrapF(webusers.Login))

	r.GET("/settings", gin.WrapF(webusers.Settings))

	log.Fatal(http.ListenAndServe(":3000", r))
}
