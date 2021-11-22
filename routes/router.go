package routes

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"

	"github.com/impact-hosting/impact-daemon/util"
)

func Launch() {
	configFileData, err := util.ReadFile(util.ConfigFile)
	if err {
		panic("Failed to read config file")
	}
	gin.SetMode("release")

	router := gin.New()
	server := socketio.NewServer()

	server.OnConnect("/", func(c socketio.Conn) error {
		if match, err := regexp.MatchString("/api/servers/[a-zA-Z0-9]+/console/?", c.URL().Path); !match || err != nil {
			return errors.New("This websocket is only available for server consoles")
		}
		id := strings.ReplaceAll(s.URL().Path, "/api/servers", "")
		id = strings.ReplaceAll(id, "console", "")
		id = strings.ReplaceAll(id, "/", "")
		if exists, err := util.ExistsContainer(id); !exists || err != nil {
			return errors.New("The specified server does not exist")
		}
		c.Join(id)
		c.SetContext(id)
		return nil
	})
	server.OnError("/", func(c socketio.Conn, e error) {

	})
	server.OnEvent("/", "command", func(c socketio.Conn, command string) {
		res := util.RunContainerCommand(c.Context(), command)
	})
	server.OnDisconnect("/", func(c socketio.Conn, reason string) {
		c.SetContext(nil)
		c.LeaveAll()
	})

	router.Use(gin.Recovery())

	// server routes
	router.GET("/api/servers")
	serve := router.Group("/api/servers/:server", Authorization(), ServerExists())
	{
		serve.GET("", GetServer)
		serve.POST("", PostServer)
		serve.PUT("", PutServer)
		serve.DELETE("", DeleteServer)

		backup := serve.Group("/backups")
		{
			backup.GET("", GetServerBackup)
			backup.POST("", PostServerBackup)
			backup.PUT("", PutServerBackup)
			backup.DELETE("", DeleteServerBackup)
		}

		console := serve.Group("/console")
		{
			console.GET("", gin.WrapH(server))
			console.POST("", gin.WrapH(server))
		}
	}

	if err := server.Serve(); err != nil {
		panic("Failed to start websocket server")
	}
	if err := router.Run(fmt.Sprintf("%s:%d", configFileData.Hostname, configFileData.Port)); err != nil {
		panic("Failed to start webserver")
	}
}
