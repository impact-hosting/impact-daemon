package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/impact-hosting/impact-daemon/structs"
)

func GetServer(c *gin.Context) *structs.Server {
	server := ExtractServer(c)
}
