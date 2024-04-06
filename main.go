package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	speech "cloud.google.com/go/speech/apiv1"
	"github.com/fabioelizandro/speech-to-text/assert"
	"github.com/fabioelizandro/speech-to-text/store"
	"github.com/fabioelizandro/speech-to-text/stt"
	"github.com/fabioelizandro/speech-to-text/web"
)

func main() {
	file := flag.String("f", "", "Audio file must be a GSC URI encoded with FLAC or WAV mono.")
	webModeOn := flag.Bool("w", false, "Webserver mode")
	flag.Parse()

	if *webModeOn {
		fmt.Println("Web mode on...")

		server := web.New("./templates", store.NewInMemoryAudioStore())
		assert.NoErr(server.Run("localhost:8080"))
		return
	}

	speechToText := stt.NewCachedSpeechToText(
		stt.NewGCSpeechToText(
			assert.Must(speech.NewClient(context.Background())),
		),
		ensureCacheDir(),
	)

	transcript := assert.Must(speechToText.Transcript(stt.NewAudioSource(*file)))
	fmt.Printf("\n\n%s\n", transcript.SpeakerDiarization().String())
}

func ensureCacheDir() string {
	userCacheDir := assert.Must(os.UserCacheDir())
	appCacheDir := filepath.Join(userCacheDir, "stt-cli")

	if _, err := os.Stat(appCacheDir); os.IsNotExist(err) {
		assert.NoErr(os.Mkdir(appCacheDir, 0754))
	}

	return appCacheDir
}
