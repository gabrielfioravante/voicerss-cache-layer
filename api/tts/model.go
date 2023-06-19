package tts

import (
	"gorm.io/gorm"
	. "voicerrs-cache-layer/db"
)

type Tts struct {
	gorm.Model
	Url   string `gorm:"uniqueIndex"`
	Audio string `gorm:"type:longtext"`
}

func (m Tts) FindAudio(url string) (string, error) {
	var tts Tts

	result := DB.Where("url = ?", url).First(&tts)

	if result.Error != nil {
		return "", result.Error
	}

	return tts.Audio, nil
}

func (m Tts) Create(url string, audio string) error {
	tts := Tts{
		Url:   url,
		Audio: audio,
	}

	result := DB.Create(&tts)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
