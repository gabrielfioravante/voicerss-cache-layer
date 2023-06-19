package config

import (
	"os"
)

type ApiConfig struct {
	VoiceRssKey string
	ServerMode  string
}

var Api *ApiConfig

func Init() {
	voiceRssKey, exists := os.LookupEnv("VOICERSS_API_KEY")

	if !exists {
		panic("VOICERSS_API_KEY is not set inside .env")
	}

	serverMode, exists := os.LookupEnv("SERVER_MODE")

	if !exists {
		serverMode = "release"
	}

	Api = &ApiConfig{
		VoiceRssKey: voiceRssKey,
		ServerMode:  serverMode,
	}
}
