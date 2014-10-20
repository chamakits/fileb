package fileb

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Fileb is the "File Battery" type, which contains many common functions needed
// for file handling.
type Fileb struct {
	innerPath string
}

// NewFileb creates a new fileb struct.
// This DOES NOT actually create a new file.
func NewFileb(path string) (*Fileb, error) {
	absFilePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	return &Fileb{
		innerPath: absFilePath,
	}, nil
}

// Path returns the path related to this file
func (filebSelf *Fileb) Path() string {
	return filebSelf.innerPath
}

func (filebSelf *Fileb) CreateFile() (*os.File, error) {
	return os.Create(filebSelf.innerPath)
}

// TODO Make this better for creating with permissions.  A nice enum type
// or something.
func (filebSelf *Fileb) CreateDir() error {
	return os.MkdirAll(filebSelf.innerPath, 0777)
}

// GetFileInfo returns the file info struct related to this file
func (filebSelf *Fileb) GetFileInfo() (os.FileInfo, error) {
	fileInfo, err := os.Stat(filebSelf.innerPath)
	if os.IsNotExist(err) {
		return nil, errors.New("File does not exist.")
	}

	if err != nil {
		return nil, err
	}

	return fileInfo, nil

}

// IsExecutable checks if the related file is executable or not.
func (filebSelf *Fileb) IsExecutable() bool {
	fileInfo, err := filebSelf.GetFileInfo()
	if err != nil {
		return false
	}
	mode := fileInfo.Mode().Perm()
	//TODO UNTESTED.  Test this.
	return mode|0111 == 0
}

// IsExecutableByCurrentUser verifies if the realted file is executable by current
// user.
func (filebSelf *Fileb) IsExecutableByCurrentUser() bool {
	//TODO this will need to be a mix of these:
	// https://groups.google.com/forum/#!topic/golang-nuts/ywS7xQYJkHY
	// http://golang.org/pkg/os/#Geteuid
	panic("Method not yet implemented")
}

// IsDirectory returns true if the related file is a directory.
func (filebSelf *Fileb) IsDirectory() (bool, error) {
	fileInfo, err := filebSelf.GetFileInfo()
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

// FileWatcher is a type of function that expects a 'Fileb' and returns nothing.
type FileWatcher func(*Fileb)

// WatchFile is a function that wathces a file.
func (filebSelf *Fileb) WatchFile(fileWatcher FileWatcher) error {
	isDirectory, err := filebSelf.IsDirectory()
	if err != nil {
		return fmt.Errorf("Problem watching file:%v\n", err)
	} else if isDirectory {
		return errors.New("Given file is a directory.  Use 'WatchDirectory'.")
	}
	panic("Method not yet implemented")
}

// WatchDirectory is a function that watches a directory.  It watches all file
// in directory, and possibly recursively.
// You can specify how many directories depth to watch files
// To not search recursively use '0'
// To search recursively infinitely use '-1'
func (filebSelf *Fileb) WatchDirectory(dept int, fileWatcher FileWatcher) error {
	isDirectory, err := filebSelf.IsDirectory()

	if err != nil {
		return fmt.Errorf("Problem watching directory:%v\n", err)
	} else if !isDirectory {
		return errors.New("Given file is not a directory.  Use 'WatchFile'.")
	}
	panic("Method not yet implemented")
}

// ReadBytes returns bytes for the related file.
func (filebSelf *Fileb) ReadBytes() ([]byte, error) {
	return ioutil.ReadFile(filebSelf.innerPath)
}
