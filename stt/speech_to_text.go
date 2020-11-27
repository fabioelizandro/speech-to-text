package stt

import (
	"context"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

type SpeechToText struct {
	client *speech.Client
}

func NewSpeechToText(client *speech.Client) SpeechToText {
	return SpeechToText{client: client}
}

func (s SpeechToText) Transcript(audio AudioSource) (Transcript, error) {
	req := &speechpb.LongRunningRecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			LanguageCode:               "pt-BR",
			EnableWordTimeOffsets:      true,
			EnableAutomaticPunctuation: true,
			MaxAlternatives:            1,
			DiarizationConfig: &speechpb.SpeakerDiarizationConfig{
				EnableSpeakerDiarization: true,
				MinSpeakerCount:          2,
				MaxSpeakerCount:          2,
			},
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
