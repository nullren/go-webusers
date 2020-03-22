package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/nullren/go-webusers/pkg/user"
	"github.com/nullren/go-webusers/pkg/webusers"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})

	users := user.New(user.NewLocalStorage())
	handlers := webusers.Handlers{Users: users}

	r.GET("/signup", gin.WrapF(handlers.SignUp))
	r.POST("/signup", gin.WrapF(handlers.SignUp))

	r.GET("/login", gin.WrapF(handlers.Login))
	r.POST("/login", gin.WrapF(handlers.Login))

	r.GET("/settings", gin.WrapF(handlers.Settings))

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(addr, r))
}
