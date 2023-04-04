package archive

import (
	"bytes"
	"io"
	"time"
)

type Archive interface {
	Create(out io.Writer, files []*File) (*Meta, error)
	Extract(in io.Reader) ([]*File, error)
}

type Meta struct {
	Size int64
}

type File struct {
	Name    string
	Size    int64
	Mode    int64
	ModTime time.Time
	Data    *bytes.Buffer
}
