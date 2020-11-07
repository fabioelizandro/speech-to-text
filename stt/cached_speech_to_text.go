package stt

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type CachedSpeechToText struct {
	underlyingImpl SpeechToText
	cacheDir       string
}

func NewCachedSpeechToText(underlyingImpl SpeechToText, cacheDir string) *CachedSpeechToText {
	return &CachedSpeechToText{underlyingImpl: underlyingImpl, cacheDir: cacheDir}
}

func (t *CachedSpeechToText) Transcript(audio AudioSource) (Transcript, error) {
	audioCacheKey := filepath.Join(
		t.cacheDir,
		base64.StdEncoding.EncodeToString([]byte(audio.URI())),
	)

	transcriptBytes, err := ioutil.ReadFile(audioCacheKey)
	if err != nil {
		transcript, err := t.underlyingImpl.Transcript(audio)
		if err != nil {
			return transcript, err
		}

		transcriptBytes, err = json.Marshal(transcript)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(audioCacheKey, transcriptBytes, 0754)
		if err != nil {
			panic(err)
		}
	}

	transcript := Transcript{}
	err = json.Unmarshal(transcriptBytes, &transcript)
	if err != nil {
		panic(err)
	}

	return transcript, nil
}
