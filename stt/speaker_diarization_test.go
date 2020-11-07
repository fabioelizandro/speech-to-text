package stt_test

import (
	"testing"

	"github.com/fabioelizandro/speech-to-text/stt"
	"github.com/stretchr/testify/assert"
)

func Test_tag_quotes_with_speaker_name(t *testing.T) {
	sd := stt.NewSpeakerDiarization()
	sd.AddWord("speaker1", "hello")

	expected := `speaker1: hello`
	assert.Equal(t, expected, sd.String())
}

func Test_appends_words_to_speaker_quote(t *testing.T) {
	sd := stt.NewSpeakerDiarization()
	sd.AddWord("speaker1", "hello")
	sd.AddWord("speaker1", "there")

	expected := `speaker1: hello there`
	assert.Equal(t, expected, sd.String())
}

func Test_breaks_line_for_new_speaker(t *testing.T) {
	sd := stt.NewSpeakerDiarization()
	sd.AddWord("speaker1", "how")
	sd.AddWord("speaker1", "are")
	sd.AddWord("speaker1", "you?")

	sd.AddWord("speaker2", "I'm")
	sd.AddWord("speaker2", "good")
	sd.AddWord("speaker2", "thanks.")

	expected := `speaker1: how are you?

speaker2: I'm good thanks.`
	assert.Equal(t, expected, sd.String())
}
