package file

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

type File struct {
	Name string
	Date []byte
	Id uuid.UUID
}

func NewFile(filename string, date []byte) (*File, error) {
	
	fileID, err := uuid.NewUUID()
	if(err != nil){
		log.Fatal("Cant generate uuid for file")
		return &File{}, err
	}

	return &File{
		Name: filename,
		Date: date,
		Id: fileID,
	}, nil
}

func (f *File) String() string {
	return fmt.Sprintf("file whith name %s, uuid: %v", f.Name, f.Id)
}