package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReturnMessage(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"message": message,
		"status":  status,
	})
}

func ReturnError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"error":  message,
		"status": status,
	})
}

func Return400(c *gin.Context) {
	ReturnError(c, http.StatusBadRequest, "Bad Request")
}

func Return500(c *gin.Context) {
	ReturnError(c, http.StatusInternalServerError, "Internal Server Error")
}
