package stt_test

import (
	"encoding/json"
	"testing"

	"github.com/fabioelizandro/speech-to-text/stt"
	"github.com/stretchr/testify/assert"
)

func Test_transcribes_using_speaker_diarization(t *testing.T) {
	transcript := stt.Transcript{}
	assert.NoError(t, json.Unmarshal([]byte(transcriptExample), &transcript))

	expectedDiarization := stt.NewSpeakerDiarization()
	expectedDiarization.AddWord("speaker1", "Hello")
	expectedDiarization.AddWord("speaker1", "there")
	expectedDiarization.AddWord("speaker2", "How")
	expectedDiarization.AddWord("speaker2", "you?")
	expectedDiarization.AddWord("speaker1", "I'm")
	expectedDiarization.AddWord("speaker1", "good")
	expectedDiarization.AddWord("speaker2", "Lucky")
	expectedDiarization.AddWord("speaker2", "you")

	assert.Equal(t, expectedDiarization, transcript.SpeakerDiarization())
}

const transcriptExample = `
[
  {
    "alternatives": [
      {
        "transcript": "Some sentence to be ignored.",
        "confidence": 0.9225482,
        "words": []
      }
    ]
  },
  {
    "alternatives": [
      {
        "words": [
          {
            "start_time": {
              "seconds": 67,
              "nanos": 600000000
            },
            "end_time": {
              "seconds": 67,
              "nanos": 900000000
            },
            "word": "Hello",
            "speaker_tag": 1
          },
          {
            "start_time": {
              "seconds": 67,
              "nanos": 900000000
            },
            "end_time": {
              "seconds": 68,
              "nanos": 100000000
            },
            "word": "there",
            "speaker_tag": 1
          },
          {
            "start_time": {
              "seconds": 69
            },
            "end_time": {
              "seconds": 71,
              "nanos": 100000000
            },
            "word": "How",
            "speaker_tag": 2
          },
          {
            "start_time": {
              "seconds": 71,
              "nanos": 100000000
            },
            "end_time": {
              "seconds": 71,
              "nanos": 400000000
            },
            "word": "you?",
            "speaker_tag": 2
          },
          {
            "start_time": {
              "seconds": 71,
              "nanos": 900000000
            },
            "end_time": {
              "seconds": 76,
              "nanos": 200000000
            },
            "word": "I'm",
            "speaker_tag": 1
          },
          {
            "start_time": {
              "seconds": 76,
              "nanos": 200000000
            },
            "end_time": {
              "seconds": 77,
              "nanos": 100000000
            },
            "word": "good",
            "speaker_tag": 1
          },
          {
            "start_time": {
              "seconds": 83,
              "nanos": 800000000
            },
            "end_time": {
              "seconds": 86,
              "nanos": 400000000
            },
            "word": "Lucky",
            "speaker_tag": 2
          },
          {
            "start_time": {
              "seconds": 86,
              "nanos": 400000000
            },
            "end_time": {
              "seconds": 86,
              "nanos": 500000000
            },
            "word": "you",
            "speaker_tag": 2
          }
        ]
      }
    ]
  }
]
`
