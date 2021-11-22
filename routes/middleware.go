package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/impact-hosting/impact-daemon/structs"
)

func ExtractServer(c *gin.Context) *structs.Server {
	server, exists := c.Get("server")
	if exists {
		return server.(*structs.Server)
	}
	panic("Could not extract server: It does not exist")
}

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func ServerExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
