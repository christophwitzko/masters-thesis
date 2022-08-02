package output

import (
	"io"
	"os"
)

type File struct {
	osFile *os.File
}

func newFileWriterFromPath(outputFile string) (io.WriteCloser, error) {
	if outputFile == "-" {
		return &File{osFile: os.Stdout}, nil
	}
	outFile, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &File{osFile: outFile}, nil
}

func (f *File) Write(p []byte) (n int, err error) {
	return f.osFile.Write(p)
}

func (f *File) Close() error {
	_ = f.osFile.Sync()
	return f.osFile.Close()
}
