package file

import "os"

type FolderInterface interface {
	GetPath() string
	Save(fileName string, file *os.File) (string, error)
	Read(fileName string) []byte
	Delete(fileName string) bool
	Exist(fileName string) bool
}

type Folder struct {
	FolderInterface
	path string
}

func NewFolder() *Folder {
	folder := Folder{}
	return &folder
}

func (f Folder) GetPath() string {
	return f.path
}
