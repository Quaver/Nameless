package main

import (
	"fmt"
	handlers2 "github.com/Swan/Nameless/handlers"
	scores2 "github.com/Swan/Nameless/handlers/scores"
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

	router.POST("/submit", scores2.Handler{}.SubmitPOST)

	router.NoRoute(func(c *gin.Context) {
		handlers2.ReturnError(c, http.StatusNotFound, "Not Found")
		return
	})

	logger.Info(fmt.Sprintf("Listening and serving HTTP on :%v", port))
	log.Fatal(router.Run(fmt.Sprintf(":%v", port)))
}
