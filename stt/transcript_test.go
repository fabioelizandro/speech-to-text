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

	expected := `speaker1: Hello there

speaker2: How you?

speaker1: I'm good

speaker2: Lucky you`
	assert.Equal(t, expected, transcript.String())
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
