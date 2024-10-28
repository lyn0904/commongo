package file

import (
	"os"
	"path/filepath"
)

type FolderInterface interface {
	GetPath() string
	Save(fileName string, data []byte, additional bool) (string, error)
	Read(fileName string) []byte
	ReadString(fileName string) string
	Delete(fileName string) bool
	Exist(fileName string) bool
}

type Folder struct {
	FolderInterface
	path string
}

func NewFolder(path string) Folder {
	folder := Folder{
		path: path,
	}
	return folder
}

func (f *Folder) GetPath() string {
	abs, err := filepath.Abs(f.path)
	if err != nil {
		return ""
	}
	return abs
}

func (f *Folder) Save(fileName string, data []byte, additional bool) (string, error) {
	abs := filepath.Join(f.GetPath(), fileName)
	var flag int
	if additional {
		flag = os.O_APPEND | os.O_RDWR
	} else {
		flag = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	}
	openFile, err := os.OpenFile(abs, flag, os.ModeSetuid)
	if err != nil {
		return "", err
	}
	defer openFile.Close()
	_, err = openFile.Write(data)
	if err != nil {
		return "", err
	}
	return abs, nil
}

func (f *Folder) Read(fileName string) []byte {
	abs := filepath.Join(f.GetPath(), fileName)
	data, err := os.ReadFile(abs)
	if err != nil {
		return nil
	}
	return data
}

func (f *Folder) ReadString(fileName string) string {
	bytes := f.Read(fileName)
	return string(bytes)
}

func (f *Folder) Delete(fileName string) bool {
	abs := filepath.Join(f.GetPath(), fileName)
	err := os.Remove(abs)
	if err != nil {
		return false
	}
	return true
}

func (f *Folder) Exist(fileName string) bool {
	abs := filepath.Join(f.GetPath(), fileName)
	_, err := os.Stat(abs)
	if err != nil {
		return false
	}
	return true
}
