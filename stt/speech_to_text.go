package stt

import (
	"context"
	"log"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

type SpeechToText struct {
	client *speech.Client
}

func NewSpeechToText() SpeechToText {
	client, err := speech.NewClient(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return SpeechToText{client: client}
}

func (s SpeechToText) Transcript(audio AudioSource) (Transcript, error) {
	req := &speechpb.LongRunningRecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			LanguageCode: "pt-BR",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Uri{Uri: audio.URI()},
		},
	}

	op, err := s.client.LongRunningRecognize(context.Background(), req)
	if err != nil {
		return Transcript{}, err
	}
	resp, err := op.Wait(context.Background())
	if err != nil {
		return Transcript{}, err
	}

	return NewTranscript(resp.Results), nil
}
