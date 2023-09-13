// file.go

import "github.com/spf13/afero"

// File represents an open file
type File struct {
  // Path, contents, etc
}

// Filesystem is an abstraction layer for file io 
type Filesystem struct {
  fs afero.Fs
}

// ReadFile reads a file from disk
func (f *Filesystem) ReadFile(path string) (*File, error) {
  // Implement reading file   
}

// WriteFile writes a file to disk
func (f *Filesystem) WriteFile(file *File) error {
  // Implement writing to disk
}
