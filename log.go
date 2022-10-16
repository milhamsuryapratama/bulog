package bulog

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
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

func (l *Logger) newEvent(level Level, done func(string)) *Event {
	if done != nil {
		done("")
		return nil
	}

	e := newEvent(l.w, level)
	if level != NoLevel {
		e.Str(LevelFieldName, level.String())
	}
	return e
}

func (l *Logger) Info() *Event {
	return l.newEvent(InfoLevel, nil)
}

func (l *Logger) Debug() *Event {
	return l.newEvent(DebugLevel, nil)
}

func (l *Logger) Warn() *Event {
	return l.newEvent(WarnLevel, nil)
}

func (l *Logger) Error() *Event {
	return l.newEvent(ErrorLevel, nil)
}

func (l *Logger) Fatal() *Event {
	return l.newEvent(FatalLevel, func(msg string) { os.Exit(1) })
}

func (l *Logger) Panic() *Event {
	return l.newEvent(PanicLevel, func(msg string) { panic(msg) })
}

// default log level Print is debug
func (l *Logger) Print(v ...interface{}) {
	l.Debug().Msg(fmt.Sprint(v...))
}

// default log level Printf is debug
func (l *Logger) Printf(format string, v ...interface{}) {
	l.Debug().Msg(fmt.Sprintf(format, v...))
}
