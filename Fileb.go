package fileb

import (
	"errors"
	"os"
)

type Fileb struct {
	Path string
}

func NewFileb(path string) *Fileb {
	return &Fileb{
		Path: path,
	}
}

func (filebSelf *Fileb) IsExecutable() bool {
	panic("Method not yet implemented")
}

func (filebSelf *Fileb) IsExecutableByCurrentUser() bool {
	panic("Method not yet implemented")
}

func (filebSelf *Fileb) IsDirectory() bool {
	panic("Method not yet implemented")
}

func (filebSelf *Fileb) GetFileInfo() os.FileInfo {
	panic("Method not yet implemented")
}

type FileWatcher func(*Fileb)

func (filebSelf *Fileb) WatchFile(fileWatcher FileWatcher) error {
	if filebSelf.IsDirectory() {
		return errors.New("Given file is a directory.  Use 'WatchDirectory'.")
	}
	panic("Method not yet implemented")
}

func (filebSelf *Fileb) WatchDirectory(recursive bool, fileWatcher FileWatcher) error {
	if !filebSelf.IsDirectory() {
		return errors.New("Given file is not a directory.  Use 'WatchFile'.")
	}
	panic("Method not yet implemented")
}
