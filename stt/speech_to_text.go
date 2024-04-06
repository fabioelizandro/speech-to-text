package stt

import (
	"fmt"
	"strings"

	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

type SpeechToText interface {
	Transcript(audio AudioSource) (Transcript, error)
}

type FakeSpeechToText struct {
}

func NewFakeSpeechToText() *FakeSpeechToText {
	return &FakeSpeechToText{}
}

func (t *FakeSpeechToText) Transcript(audio AudioSource) (Transcript, error) {
	transcriptText := fmt.Sprintf("%s: Simple text", audio.URI())

	wordInfos := []*speechpb.WordInfo{}
	for _, word := range strings.Split(transcriptText, " ") {
		wordInfos = append(wordInfos, &speechpb.WordInfo{
			Word: word,
		})
	}

	return NewTranscript([]*speechpb.SpeechRecognitionResult{
		{
			Alternatives: []*speechpb.SpeechRecognitionAlternative{
				{
					Transcript: transcriptText,
					Confidence: 0.95,
					Words:      wordInfos,
				},
			},
		},
	}), nil
}
