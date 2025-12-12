package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// BasicAuthMiddleware checks Authorization header Basic user:pass
func BasicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, p, ok := c.Request.BasicAuth()
		adminUser := os.Getenv("ADMIN_USER")
		adminPass := os.Getenv("ADMIN_PASS")

		if !ok || u != adminUser || p != adminPass {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Next()
	}
}
