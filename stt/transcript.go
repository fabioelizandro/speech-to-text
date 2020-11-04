package stt

import (
	"fmt"
	"io/ioutil"
	"strings"

	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

type Transcript struct {
	results []*speechpb.SpeechRecognitionResult
}

func NewTranscript(results []*speechpb.SpeechRecognitionResult) Transcript {
	return Transcript{results: results}
}

func (t Transcript) Report(filename string) error {
	return ioutil.WriteFile(filename, []byte(t.content()), 644)
}

func (t Transcript) Print() {
	fmt.Print(t.content())
}

func (t Transcript) content() string {
	content := strings.Builder{}
	for _, result := range t.results {
		content.WriteString(result.Alternatives[0].Transcript)
		content.WriteString("\n")
	}

	return content.String()
}
