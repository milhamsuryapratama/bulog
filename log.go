package bulog

import (
	"io"
	"io/ioutil"
	"strconv"
)

// Level defines log levels.
type Level int8

const (
	// DebugLevel defines debug log level.
	DebugLevel Level = iota
	// InfoLevel defines info log level.
	InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel
	// FatalLevel defines fatal log level.
	FatalLevel
	// PanicLevel defines panic log level.
	PanicLevel
	// NoLevel defines an absent log level.
	NoLevel
	// Disabled disables the logger.
	Disabled

	// TraceLevel defines trace log level.
	TraceLevel Level = -1
	// Values less than TraceLevel are handled as numbers.
)

func (l Level) String() string {
	switch l {
	case TraceLevel:
		return LevelTraceValue
	case DebugLevel:
		return LevelDebugValue
	case InfoLevel:
		return LevelInfoValue
	case WarnLevel:
		return LevelWarnValue
	case ErrorLevel:
		return LevelErrorValue
	case FatalLevel:
		return LevelFatalValue
	case PanicLevel:
		return LevelPanicValue
	case Disabled:
		return "disabled"
	case NoLevel:
		return ""
	}
	return strconv.Itoa(int(l))
}

type Logger struct {
	w     LevelWriter
	level Level
}

func New(w io.Writer) Logger {
	if w == nil {
		w = ioutil.Discard
	}
	lw, ok := w.(LevelWriter)
	if !ok {
		lw = levelWriterAdapter{w}
	}
	return Logger{w: lw, level: TraceLevel}
}

func (l *Logger) newEvent(level Level) *Event {
	e := newEvent(l.w, level)
	if level != NoLevel {
		e.Str(LevelFieldName, level.String())
	}
	return e
}

func (l *Logger) Info() *Event {
	return l.newEvent(InfoLevel)
}

func (l *Logger) Debug() *Event {
	return l.newEvent(DebugLevel)
}

func (l *Logger) Warn() *Event {
	return l.newEvent(WarnLevel)
}

func (l *Logger) Error() *Event {
	return l.newEvent(ErrorLevel)
}

func (l *Logger) Fatal() *Event {
	return l.newEvent(ErrorLevel)
}
