// Package zip make a packing base on raw zip package.
// It let you easy to pack files and folds directly.
package zip

import (
	"archive/zip"
	"bytes"
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
		return w.WriteFile(path)
	}); err != nil {
		return err
	}
	return nil
}

// WriteFile write a file to zip.
func (w *Writer) WriteFile(path string) error {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return w.WriteBytes(path, fileBytes)
}

// WriteBytes write file bytes to zip.
func (w *Writer) WriteBytes(path string, fileBytes []byte) error {
	writer, err := w.Create(w.Prefix + path)
	if err != nil {
		w.Close()
		return err
	}
	if _, err = writer.Write(fileBytes); err != nil {
		return err
	}
	return nil
}

// Pack src into dst
func Pack(dst, src string) error {
	buf := bytes.Buffer{}

	w := NewWriter(&buf)
	if err := w.WriteFiles(src); err != nil {
		return err
	}
	w.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	if _, err := buf.WriteTo(dstFile); err != nil {
		return err
	}
	dstFile.Close()

	return nil
}
