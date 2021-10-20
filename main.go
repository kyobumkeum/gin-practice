package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := createServer()
	server.ListenAndServe()
}

func createServer() *http.Server {
	return &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        createRouter(),
		ReadTimeout:    0,
		WriteTimeout:   0,
		MaxHeaderBytes: 1 << 20,
	}
}

func createRouter() *gin.Engine {
	r := gin.New()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
