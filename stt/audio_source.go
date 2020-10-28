package stt

type AudioSource struct {
	uri string
}

func NewAudioSource(uri string) AudioSource {
	return AudioSource{uri: uri}
}

func (s AudioSource) URI() string {
	return s.uri
}
