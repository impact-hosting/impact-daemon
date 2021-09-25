package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/impact-hosting/impact-daemon/util"
)

var ()

func Launch() {
	configFileData, err := util.ReadFile(util.ConfigFile)
	gin.SetMode("release")

	router := gin.New()
	router.Use((gin.Recovery()))

	// server routes
	router.GET("/api/servers")

	router.Run(fmt.Sprintf("%s:%d", configFileData, configFileData))
}
