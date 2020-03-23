package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/nullren/go-webusers/pkg/session"
	memstore2 "github.com/nullren/go-webusers/pkg/session/memstore"
	"github.com/nullren/go-webusers/pkg/user"
	"github.com/nullren/go-webusers/pkg/user/memstore"
	"github.com/nullren/go-webusers/pkg/webusers"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})

	handlers := webusers.Handlers{
		Users:    user.New(memstore.NewStore()),
		Sessions: session.New(memstore2.NewStore()),
	}

	r.GET("/signup", gin.WrapF(handlers.SignUp))
	r.POST("/signup", gin.WrapF(handlers.SignUp))

	r.GET("/login", gin.WrapF(handlers.Login))
	r.POST("/login", gin.WrapF(handlers.Login))

	{
		basicAuth := r.Group("/")
		basicAuth.Use(AuthenticationRequired(handlers))
		basicAuth.GET("/settings", gin.WrapF(handlers.Settings))
	}

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(addr, r))
}

func AuthenticationRequired(h webusers.Handlers) func(c *gin.Context) {
	return func(c *gin.Context) {
		if u := h.Sessions.Read(c.Request); u == nil {
			c.String(401, "unauthorized :(")
			c.Abort()
			return
		}
		c.Next()
	}
}
