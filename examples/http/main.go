package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/nullren/go-webusers"
)

var port = "3000"

func init() {
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})

	webusers.AddWebUsers(r)

	addr := fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(addr, r))
}
