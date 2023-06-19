package server

import (
	"voicerrs-cache-layer/api/config"
	"voicerrs-cache-layer/api/tts"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	gin.SetMode(config.Api.ServerMode)

	r := gin.Default()
	r.Use(cors.Default())

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/tts", tts.Get)
	}

	r.Run()
}
