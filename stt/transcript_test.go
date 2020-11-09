package stt_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/fabioelizandro/speech-to-text/stt"
	"github.com/stretchr/testify/assert"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func Test_transcribes_using_speaker_diarization(t *testing.T) {
	transcript := stt.Transcript{}
	assert.NoError(t, json.Unmarshal([]byte(transcriptExample), &transcript))

	expectedDiarization := stt.NewSpeakerDiarization()
	expectedDiarization.AddWord("speaker1", "Hello", 67*time.Second)
	expectedDiarization.AddWord("speaker1", "there", 67*time.Second)
	expectedDiarization.AddWord("speaker2", "How", 69*time.Second)
	expectedDiarization.AddWord("speaker2", "you?", 71*time.Second)
	expectedDiarization.AddWord("speaker1", "I'm", 71*time.Second)
	expectedDiarization.AddWord("speaker1", "good", 76*time.Second)
	expectedDiarization.AddWord("speaker2", "Lucky", 83*time.Second)
	expectedDiarization.AddWord("speaker2", "you", 86*time.Second)

	assert.Equal(t, expectedDiarization, transcript.SpeakerDiarization())
}

func Test_returns_empty_speaker_diarization_when_results_are_empty(t *testing.T) {
	transcript := stt.NewTranscript([]*speechpb.SpeechRecognitionResult{})

	expectedDiarization := stt.NewSpeakerDiarization()

	assert.Equal(t, expectedDiarization, transcript.SpeakerDiarization())
}

func Test_transcribes_without_speaker_diarization(t *testing.T) {
	transcript := stt.Transcript{}
	assert.NoError(t, json.Unmarshal([]byte(transcriptExample), &transcript))

	expected := `[0s] Hello there How you?

[5s] I'm good Lucky you`
	assert.Equal(t, expected, transcript.String())
}

func Test_returns_empty_string_for_empty_results(t *testing.T) {
	transcript := stt.NewTranscript([]*speechpb.SpeechRecognitionResult{})

	assert.Equal(t, "", transcript.String())
}

const transcriptExample = `
[
  {
    "alternatives": [
      {
        "transcript": "Hello there How you?",
        "confidence": 0.9225482,
        "words": [
          {
            "start_time": {},
            "end_time": {
              "seconds": 3
            },
            "word": "Hello"
          },
          {
            "start_time": {
              "seconds": 3
            },
            "end_time": {
              "seconds": 3
            },
            "word": "there"
          },
          {
            "start_time": {
              "seconds": 3
            },
            "end_time": {
              "seconds": 4
            },
            "word": "How"
          },
          {
            "start_time": {
              "seconds": 4
            },
            "end_time": {
              "seconds": 5
            },
            "word": "you?"
          }
        ]
      }
    ]
  },
  {
    "alternatives": [
      {
        "transcript": "I'm good Lucky you",
        "confidence": 0.9115471,
        "words": [
          {
            "start_time": {
              "seconds": 5,
              "nanos": 500000000
            },
            "end_time": {
              "seconds": 5
            },
            "word": "I'm"
          },
          {
            "start_time": {
              "seconds": 6
            },
            "end_time": {
              "seconds": 6
            },
            "word": "good"
          },
          {
            "start_time": {
              "seconds": 6
            },
            "end_time": {
              "seconds": 7
            },
            "word": "Lucky"
          },
          {
            "start_time": {
              "seconds": 7
            },
            "end_time": {
              "seconds": 8
            },
            "word": "you"
          }
        ]
      }
    ]
  },
  {
    "alternatives": [
      {
        "words": [
          {
            "start_time": {
              "seconds": 67
            },
            "end_time": {
              "seconds": 67
            },
            "word": "Hello",
            "speaker_tag": 1
          },
          {
            "start_time": {
              "seconds": 67
            },
            "end_time": {
              "seconds": 68
            },
            "word": "there",
            "speaker_tag": 1
          },
          {
            "start_time": {
              "seconds": 69
            },
            "end_time": {
              "seconds": 71
            },
            "word": "How",
            "speaker_tag": 2
          },
          {
            "start_time": {
              "seconds": 71
            },
            "end_time": {
              "seconds": 71
            },
            "word": "you?",
            "speaker_tag": 2
          },
          {
            "start_time": {
              "seconds": 71
            },
            "end_time": {
              "seconds": 76
            },
            "word": "I'm",
            "speaker_tag": 1
          },
          {
            "start_time": {
              "seconds": 76
            },
            "end_time": {
              "seconds": 77
            },
            "word": "good",
            "speaker_tag": 1
          },
          {
            "start_time": {
              "seconds": 83
            },
            "end_time": {
              "seconds": 86
            },
            "word": "Lucky",
            "speaker_tag": 2
          },
          {
            "start_time": {
              "seconds": 86
            },
            "end_time": {
              "seconds": 86
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
