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
	report := strings.Builder{}
	for _, result := range t.results {
		report.WriteString(result.Alternatives[0].Transcript)
		report.WriteString("\n")
	}

	return ioutil.WriteFile(filename, []byte(report.String()), 644)
}

func (t Transcript) Print() {
	for _, result := range t.results {
		fmt.Printf(`"%s"\n`, result.Alternatives[0].Transcript)
	}
}
