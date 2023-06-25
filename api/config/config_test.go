package config

import (
    "testing"
	"github.com/stretchr/testify/assert"
    "os"
)

func TestInit(t *testing.T) {
    os.Setenv("VOICERSS_API_KEY", "test")
    os.Setenv("PORT", "5000")

    Init()
    assert.NotNil(t, Api)
    assert.Equal(t, Api.VoiceRssKey, "test")
    assert.Equal(t, Api.Port, "5000")
}
