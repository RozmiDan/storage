package storage

import (
	"fmt"

	"github.com/RozmiDan/storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	сollection map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		сollection: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, date []byte) (*file.File, error) {
	newf, err := file.NewFile(filename, date)
	if err != nil {
		return &file.File{}, err
	}
	s.сollection[newf.Id] = newf

	return newf, nil
}

func (s *Storage) GetById(fileId uuid.UUID) (*file.File, error) {
	if val, ok := s.сollection[fileId]; ok {
		return val, nil
	} else {
		return nil, fmt.Errorf("file %v not found", val.Id)
	}
}