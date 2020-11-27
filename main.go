package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	speech "cloud.google.com/go/speech/apiv1"
	"github.com/fabioelizandro/speech-to-text/stt"
)

func main() {
	file := flag.String("f", "", "Audio file must be a GSC URI encoded with FLAC or WAV mono.")
	flag.Parse()

	cacheDir, err := ensureCacheDir()
	if err != nil {
		panic(err)
	}

	googleSpeech, err := speech.NewClient(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	speechToText := stt.NewCachedSpeechToText(stt.NewSpeechToText(googleSpeech), cacheDir)
	transcript, err := speechToText.Transcript(stt.NewAudioSource(*file))
	if err != nil {
		panic(err)

	}

	fmt.Printf("\n\n%s\n", transcript.SpeakerDiarization().String())
}

func ensureCacheDir() (string, error) {
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	appCacheDir := filepath.Join(userCacheDir, "stt-cli")

	if _, err := os.Stat(appCacheDir); os.IsNotExist(err) {
		err := os.Mkdir(appCacheDir, 0754)
		if err != nil {
			return "", err
		}
	}

	return appCacheDir, nil
}
