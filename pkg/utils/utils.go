package utils

import (
	"bufio"
	"os"
	"path"
)

// full stop if the file errors
func ScanFile(fileName, relativePath string) (*File, error) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filePath := path.Join(cwd, relativePath, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	return &File{
		file:   file,
		reader: bufio.NewReader(file),
	}, nil
}

type FileReader interface {
	ReadLine() (string, error)
	Close() error
}

type File struct {
	file   *os.File
	reader *bufio.Reader
}

func (f *File) ReadLine() (string, error) {
	line, _, err := f.reader.ReadLine()
	if err != nil {
		return "", err
	}

	return string(line), nil
}

func (f *File) CloseFile() error {
	return f.file.Close()
}
