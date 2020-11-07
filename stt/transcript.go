package stt

import (
	"encoding/json"
	"fmt"

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
		speakerDiarization.AddWord(fmt.Sprintf("speaker%d", word.SpeakerTag), word.Word)
	}

	return speakerDiarization
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
