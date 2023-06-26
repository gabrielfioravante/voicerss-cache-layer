package config

import (
	"os"
)

type ApiConfig struct {
	VoiceRssKey string
	Port        string
}

var Api *ApiConfig

func Init() {
	voiceRssKey, exists := os.LookupEnv("VOICERSS_API_KEY")

	if !exists {
		panic("VOICERSS_API_KEY is not set")
	}

	port, exists := os.LookupEnv("PORT")

	if !exists {
		port = "8080"
	}

	Api = &ApiConfig{
		VoiceRssKey: voiceRssKey,
		Port:        port,
	}
}
