package utils

import (
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetIpFromRequest Gets the client's proper ip address from a request.
func GetIpFromRequest(c *gin.Context) string {
	if ip := net.ParseIP(strings.TrimSpace(c.GetHeader("CF-Connecting-IP"))); ip != nil {
		return ip.String()
	}

	if ip := net.ParseIP(strings.TrimSpace(c.GetHeader("X-Real-IP"))); ip != nil {
		return ip.String()
	}

	if xff := c.GetHeader("X-Forwarded-For"); xff != "" {
		for _, part := range strings.Split(xff, ",") {
			if ip := net.ParseIP(strings.TrimSpace(part)); ip != nil {
				return ip.String()
			}
		}
	}

	host, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
	if ip := net.ParseIP(host); ip != nil {
		return ip.String()
	}

	return "::1"
}
