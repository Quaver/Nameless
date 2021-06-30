package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/Swan/Nameless/src/handlers"
	"github.com/Swan/Nameless/src/handlers/scores"
	"log"
	"net/http"
)

type server struct{}

// Start Starts the API server on a given port
func (s server) Start(port int) {
	router := gin.Default()

	router.POST("/submit", scores.Handler{}.SubmitPOST)

	router.NoRoute(func(c *gin.Context) {
		handlers.ReturnError(c, http.StatusNotFound, "Not Found")
		return
	})

	log.Fatal(router.Run(fmt.Sprintf(":%v", port)))
}
