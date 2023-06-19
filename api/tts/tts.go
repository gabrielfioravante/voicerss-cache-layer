package tts

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"voicerrs-cache-layer/api/tts/params"
)

func getAudio(p *params.Params) (string, int, error) {
	url := p.GenUrl()
	tts := Tts{}

	audio, err := tts.FindAudio(url)

	if err == nil {
		return audio, http.StatusOK, nil
	}

	audio, status, err := fetchFromVoiceRss(url)

	if err != nil {
		return "", status, err
	}

	audio = string(audio)

	if err := tts.Create(url, audio); err != nil {
		return "", http.StatusInternalServerError, err
	}

	return audio, http.StatusOK, nil
}

func fetchFromVoiceRss(url string) (string, int, error) {
	res, err := http.Get(url)

	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	defer res.Body.Close()

	audio, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	if strings.HasPrefix(string(audio), "ERROR") {
		message := strings.TrimPrefix(string(audio), "ERROR: ")

		return "", http.StatusBadRequest, errors.New("[VoiceRSS] " + message)
	}

	return string(audio), http.StatusOK, nil
}
