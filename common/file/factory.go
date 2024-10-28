package file

type Factory struct {
	folderMap map[string]Folder
}

func NewFileFactory() Factory {
	fileFactory := Factory{}
	fileFactory.folderMap = make(map[string]Folder)
	return fileFactory
}

func (f *Factory) Create(name, path string) {
	folder := NewFolder(path)
	f.folderMap[name] = folder
}

func (f *Factory) Get(name string) Folder {
	return f.folderMap[name]
}
