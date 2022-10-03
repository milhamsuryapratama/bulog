package bulog

import "io"

type LevelWriter interface {
	io.Writer
	WriteLevel(level Level, p []byte) (n int, err error)
}

type levelWriterAdapter struct {
	io.Writer
}

func (lw levelWriterAdapter) WriteLevel(level Level, p []byte) (n int, err error) {
	return lw.Write(p)
}
