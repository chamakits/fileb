package fileb

import (
	"errors"
	"os"
	"path/filepath"
)

type Fileb struct {
	Path string
}

func NewFileb(path string) (*Fileb, error) {
	absFilePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	return &Fileb{
		Path: absFilePath,
	}, nil
}

func (filebSelf *Fileb) GetFileInfo() (os.FileInfo, error) {
	fileInfo, err := os.Stat(filebSelf.Path)
	if os.IsNotExist(err) {
		return nil, errors.New("File does not exist.")
	}

	if err != nil {
		return nil, err
	}

	return fileInfo, nil

}

func (filebSelf *Fileb) IsExecutable() bool {
	fileInfo, err := filebSelf.GetFileInfo()
	if err != nil {
		return false
	}
	mode := fileInfo.Mode().Perm()
	//TODO UNTESTED.  Test this.
	return mode|0111 == 0
}

func (filebSelf *Fileb) IsExecutableByCurrentUser() bool {
	//TODO this will need to be a mix of these:
	// https://groups.google.com/forum/#!topic/golang-nuts/ywS7xQYJkHY
	// http://golang.org/pkg/os/#Geteuid
	panic("Method not yet implemented")
}

func (filebSelf *Fileb) IsDirectory() bool {
	fileInfo, err := filebSelf.GetFileInfo()
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
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
