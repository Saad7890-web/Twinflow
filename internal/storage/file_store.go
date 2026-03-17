package storage

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/Saad7890-web/twinflow/pkg/models"
)

type FileStore struct {
	file *os.File
	mu   sync.Mutex
}

func NewFileStore(path string) (*FileStore, error) {

	file, err := os.OpenFile(
		path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)

	if err != nil {
		return nil, err
	}

	return &FileStore{
		file: file,
	}, nil
}

func (s *FileStore) Save(record models.TrafficRecord) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := json.Marshal(record)
	if err != nil {
		return err
	}

	_, err = s.file.Write(append(data, '\n'))

	return err
}