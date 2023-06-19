package tts

import (
	"net/http"
	"voicerrs-cache-layer/api/config"
	"voicerrs-cache-layer/api/tts/params"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	params, err := params.Parse(c, &config.Api.VoiceRssKey)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	audio, status, err := getAudio(&params)

	if err != nil {
		c.String(status, err.Error())
	}

    c.String(status, audio)
}
