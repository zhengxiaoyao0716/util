package zip

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Writer implements a zip file writer.
type Writer struct {
	zip.Writer
	Prefix string // Path prefix of writed files.
}

// NewWriter returns a new Writer writing a zip file to w.
func NewWriter(w io.Writer) *Writer {
	return &Writer{*zip.NewWriter(w), ""}
}

// WriteFiles write all files in path to zip.
func (w *Writer) WriteFiles(path string) error {
	if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info == nil || info.IsDir() {
			return nil
		}

		fileBytes, err := ioutil.ReadFile(path)
		if err != nil {
			w.Close()
			return nil
		}
		w.WriteBytes(path, fileBytes)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// WriteFile write a file to zip.
func (w *Writer) WriteFile(path string) error {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		w.Close()
		return nil
	}
	w.WriteBytes(path, fileBytes)
	return nil
}

// WriteBytes write file bytes to zip.
func (w *Writer) WriteBytes(path string, fileBytes []byte) error {
	writer, err := w.Create(w.Prefix + path)
	if err != nil {
		return err
	}
	if _, err = writer.Write(fileBytes); err != nil {
		return err
	}
	return nil
}
