package main

import (
	"fmt"
	"github.com/Swan/Nameless/handlers"
	"github.com/Swan/Nameless/handlers/scores/chicken"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"log"
	"net/http"
)

type server struct{}

// Start Starts the API server on a given port
func (s server) Start(port int) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	router.POST("/submit", chicken.Handler{}.SubmitPOST)

	router.NoRoute(func(c *gin.Context) {
		handlers.ReturnError(c, http.StatusNotFound, "Not Found")
		return
	})

	logger.Info(fmt.Sprintf("Listening and serving HTTP on :%v", port))
	log.Fatal(router.Run(fmt.Sprintf(":%v", port)))
}
