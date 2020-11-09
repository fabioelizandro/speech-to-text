package stt

import (
	"fmt"
	"strings"
	"time"
)

type SpeakerDiarization struct {
	quotes []*speakerQuote
}

type speakerQuote struct {
	speaker  string
	quote    string
	duration time.Duration
}

func NewSpeakerDiarization() *SpeakerDiarization {
	return &SpeakerDiarization{
		quotes: []*speakerQuote{},
	}
}

func (d *SpeakerDiarization) AddWord(speaker string, word string, duration time.Duration) {
	if len(d.quotes) == 0 {
		d.quotes = append(d.quotes, &speakerQuote{speaker: speaker, quote: word, duration: duration})
		return
	}

	lastSpeakerQuote := d.quotes[len(d.quotes)-1]

	if lastSpeakerQuote.speaker != speaker {
		d.quotes = append(d.quotes, &speakerQuote{speaker: speaker, quote: word, duration: duration})
		return
	}

	lastSpeakerQuote.quote = fmt.Sprintf("%s %s", lastSpeakerQuote.quote, word)
}

func (d *SpeakerDiarization) String() string {
	content := strings.Builder{}
	for _, quote := range d.quotes {
		content.WriteString(quote.string())
		content.WriteString("\n\n")
	}

	return strings.TrimSpace(content.String())
}

func (q speakerQuote) string() string {
	return fmt.Sprintf("[%s] %s: %s", q.duration.Truncate(time.Second), q.speaker, q.quote)
}
