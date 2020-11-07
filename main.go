package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fabioelizandro/speech-to-text/stt"
)

const usage = `Usage: transcript <audiofile>

Audio file must be a GSC URI encoded with FLAC or WAV mono.

NOTE: files with the same name won't hit Google's API to avoid unwanted charges. Instead, the application will cache all
transcripts in your OS cache directory under the folder "stt-cli".
`

func main() {
	if len(os.Args) < 2 {
		_, _ = fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}

	cacheDir, err := ensureCacheDir()
	if err != nil {
		panic(err)
	}

	speechToText := stt.NewCachedSpeechToText(stt.NewSpeechToText(), cacheDir)
	transcript, err := speechToText.Transcript(stt.NewAudioSource(os.Args[1]))
	if err != nil {
		panic(err)

	}

	fmt.Printf("\n\n%s\n", transcript.String())
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
