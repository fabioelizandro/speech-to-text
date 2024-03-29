package stt

type SpeechToText interface {
	Transcript(audio AudioSource) (Transcript, error)
}
