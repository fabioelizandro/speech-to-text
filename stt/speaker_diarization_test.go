package stt_test

import (
	"testing"
	"time"

	"github.com/fabioelizandro/speech-to-text/stt"
	"github.com/stretchr/testify/assert"
)

func Test_tag_quotes_with_speaker_name(t *testing.T) {
	sd := stt.NewSpeakerDiarization()
	sd.AddWord("speaker1", "hello", 0*time.Second)

	expected := `[0s] speaker1: hello`
	assert.Equal(t, expected, sd.String())
}

func Test_appends_words_to_speaker_quote(t *testing.T) {
	sd := stt.NewSpeakerDiarization()
	sd.AddWord("speaker1", "hello", 0*time.Second)
	sd.AddWord("speaker1", "there", 0*time.Second)

	expected := `[0s] speaker1: hello there`
	assert.Equal(t, expected, sd.String())
}

func Test_breaks_line_for_new_speaker(t *testing.T) {
	sd := stt.NewSpeakerDiarization()
	sd.AddWord("speaker1", "how", 0*time.Second)
	sd.AddWord("speaker1", "are", 0*time.Second)
	sd.AddWord("speaker1", "you?", 0*time.Second)

	sd.AddWord("speaker2", "I'm", 0*time.Second)
	sd.AddWord("speaker2", "good", 0*time.Second)
	sd.AddWord("speaker2", "thanks.", 0*time.Second)

	expected := `[0s] speaker1: how are you?

[0s] speaker2: I'm good thanks.`
	assert.Equal(t, expected, sd.String())
}

func Test_adds_duration_stamp_for_each_speaker_speech(t *testing.T) {
	sd := stt.NewSpeakerDiarization()
	sd.AddWord("speaker1", "how", 1*time.Second)
	sd.AddWord("speaker1", "are", 2*time.Second)
	sd.AddWord("speaker1", "you?", 3*time.Second)

	sd.AddWord("speaker2", "I'm", 4*time.Second)
	sd.AddWord("speaker2", "good", 5*time.Second)
	sd.AddWord("speaker2", "thanks.", 6*time.Second)

	expected := `[1s] speaker1: how are you?

[4s] speaker2: I'm good thanks.`
	assert.Equal(t, expected, sd.String())
}

func Test_formats_duration_truncating_to_seconds(t *testing.T) {
	sd := stt.NewSpeakerDiarization()
	sd.AddWord("speaker1", "a", 1*time.Nanosecond)
	sd.AddWord("speaker2", "b", 1*time.Second)
	sd.AddWord("speaker1", "c", 1*time.Minute)
	sd.AddWord("speaker2", "d", 67*time.Second)
	sd.AddWord("speaker1", "e", 61*time.Minute)

	expected := `[0s] speaker1: a

[1s] speaker2: b

[1m0s] speaker1: c

[1m7s] speaker2: d

[1h1m0s] speaker1: e`
	assert.Equal(t, expected, sd.String())
}
