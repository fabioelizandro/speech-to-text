package stt

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

type Transcript struct {
	result transcriptResult
}

type transcriptResult interface {
	paragraphs() []*speechpb.SpeechRecognitionResult
	diarization() []*speechpb.WordInfo
	export() []*speechpb.SpeechRecognitionResult
}

type successfulResult struct {
	underlyingResult []*speechpb.SpeechRecognitionResult
}

type emptyResult struct {
}

func NewTranscript(results []*speechpb.SpeechRecognitionResult) Transcript {
	return Transcript{result: newTranscriptResult(results)}
}

func newTranscriptResult(results []*speechpb.SpeechRecognitionResult) transcriptResult {
	if len(results) == 0 {
		return &emptyResult{}
	}

	return &successfulResult{underlyingResult: results}
}

func (t Transcript) SpeakerDiarization() *SpeakerDiarization {
	speakerDiarization := NewSpeakerDiarization()
	speakerDiarizationWords := t.result.diarization()
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
	content := strings.Builder{}
	for _, result := range t.result.paragraphs() {
		bestAlternative := result.Alternatives[0]

		transcript := bestAlternative.Transcript
		startAt := bestAlternative.Words[0].StartTime.AsDuration().Truncate(time.Second)

		content.WriteString(fmt.Sprintf("[%s] %s", startAt, transcript))
		content.WriteString("\n\n")
	}

	return strings.TrimSpace(content.String())
}

func (t Transcript) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.result.export())
}

func (t *Transcript) UnmarshalJSON(data []byte) error {
	var results []*speechpb.SpeechRecognitionResult
	err := json.Unmarshal(data, &results)
	if err != nil {
		return err
	}

	t.result = newTranscriptResult(results)
	return nil
}

func (s successfulResult) paragraphs() []*speechpb.SpeechRecognitionResult {
	return s.underlyingResult[0 : len(s.underlyingResult)-1]
}

func (s successfulResult) diarization() []*speechpb.WordInfo {
	return s.underlyingResult[len(s.underlyingResult)-1].Alternatives[0].Words
}

func (s successfulResult) export() []*speechpb.SpeechRecognitionResult {
	return s.underlyingResult
}

func (e emptyResult) paragraphs() []*speechpb.SpeechRecognitionResult {
	return []*speechpb.SpeechRecognitionResult{}
}

func (e emptyResult) diarization() []*speechpb.WordInfo {
	return []*speechpb.WordInfo{}
}

func (e emptyResult) export() []*speechpb.SpeechRecognitionResult {
	return []*speechpb.SpeechRecognitionResult{}
}
