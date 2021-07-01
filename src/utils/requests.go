package utils

import "github.com/gin-gonic/gin"

// GetIpFromRequest Gets the client's proper ip address from a request.
func GetIpFromRequest(c *gin.Context) string {
	// Running under NGINX
	ip := c.GetHeader("X-Forwarded-For")
	
	if ip != "" {
		return ip
	}
	
	return "::1"
}
