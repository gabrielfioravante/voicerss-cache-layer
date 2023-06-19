package api

import (
	"voicerrs-cache-layer/api/tts"
	. "voicerrs-cache-layer/db"
)

func AutoMigrate() {
	DB.AutoMigrate(&tts.Tts{})
}
