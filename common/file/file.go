package file

type FileFactory struct {
	folderMap map[string]Folder
}

func NewFileFactory() FileFactory {
	fileFactory := FileFactory{}
	fileFactory.folderMap = make(map[string]Folder)
	return fileFactory
}

func (f *FileFactory) Get(name string) Folder {
	folder := f.folderMap[name]
	if folder != nil {
		return folder
	}
	return FolderImpl{}
}
