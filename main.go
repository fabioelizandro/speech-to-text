package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fabioelizandro/speech-to-text/stt"
)

const usage = `Usage: transcript <audiofile>

Audio file must be a GSC URI encoded with FLAC or WAV mono.
`

func main() {
	if len(os.Args) < 2 {
		_, _ = fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}

	speechToText := stt.NewSpeechToText()
	transcript, err := speechToText.Transcript(stt.NewAudioSource(os.Args[1]))
	if err != nil {
		log.Fatal(err)

	}

	transcript.Print()
}
