package store

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type AudioStore interface {
	SaveAudio(r io.Reader, fileName string) error
	ListSavedFiles() ([]string, error)
}

// --- Real Impl

type FSAudioStore struct {
	path string
}

func NewFSAudioStore(path string) *FSAudioStore {
	return &FSAudioStore{path: path}
}

func (store *FSAudioStore) SaveAudio(r io.Reader, fileName string) error {
	file, err := os.Create(filepath.Join(store.path, fileName))
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, r)
	return err
}

func (store *FSAudioStore) ListSavedFiles() ([]string, error) {
	var fileList []string

	files, err := ioutil.ReadDir(store.path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		fileList = append(fileList, file.Name())
	}
	return fileList, nil
}

// --- Fake Impl

type InMemoryAudioStore struct {
	audio map[string][]byte
}

func NewInMemoryAudioStore() *InMemoryAudioStore {
	return &InMemoryAudioStore{audio: make(map[string][]byte)}
}

func (store *InMemoryAudioStore) SaveAudio(r io.Reader, fileName string) error {
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, r)
	if err != nil {
		return err
	}

	store.audio[fileName] = buf.Bytes()
	return nil
}

func (store *InMemoryAudioStore) ListSavedFiles() ([]string, error) {
	var fileList []string

	for fileName := range store.audio {
		fileList = append(fileList, fileName)
	}
	return fileList, nil
}
