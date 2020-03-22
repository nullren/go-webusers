package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/nullren/go-webusers/pkg/webusers"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})

	r.GET("/signup", gin.WrapF(webusers.Signup))
	r.POST("/signup", gin.WrapF(webusers.Signup))

	r.GET("/login", gin.WrapF(webusers.Login))
	r.POST("/login", gin.WrapF(webusers.Login))

	r.GET("/settings", gin.WrapF(webusers.Settings))

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(addr, r))
}
