package stt

import (
	"encoding/json"
	"fmt"
	"strings"

	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

type Transcript struct {
	results []*speechpb.SpeechRecognitionResult
}

func NewTranscript(results []*speechpb.SpeechRecognitionResult) Transcript {
	return Transcript{results: results}
}

func (t Transcript) SpeakerDiarization() *SpeakerDiarization {
	speakerDiarization := NewSpeakerDiarization()

	if len(t.results) == 0 {
		return speakerDiarization
	}

	speakerDiarizationWords := t.results[len(t.results)-1].Alternatives[0].Words
	for _, word := range speakerDiarizationWords {
		speakerDiarization.AddWord(
			fmt.Sprintf("speaker%d", word.SpeakerTag),
			word.Word,
			word.StartTime.AsDuration(),
		)
	}

	return speakerDiarization
}

func (t Transcript) String() string {
	if len(t.results) == 0 {
		return ""
	}

	content := strings.Builder{}
	for _, result := range t.results[0 : len(t.results)-1] {
		bestAlternative := result.Alternatives[0]

		transcript := bestAlternative.Transcript
		startAt := bestAlternative.Words[0].StartTime.AsDuration()

		content.WriteString(fmt.Sprintf("[%s] %s", startAt, transcript))
		content.WriteString("\n\n")
	}

	return strings.TrimSpace(content.String())
}

func (t Transcript) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.results)
}

func (t *Transcript) UnmarshalJSON(data []byte) error {
	var results []*speechpb.SpeechRecognitionResult
	err := json.Unmarshal(data, &results)
	if err != nil {
		return err
	}

	t.results = results
	return nil
}
