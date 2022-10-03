package bulog

import "fmt"

type Event struct {
	buf   []byte
	w     LevelWriter
	level Level
}

func newEvent(w LevelWriter, level Level) *Event {
	e := &Event{}
	enc := Encoder{}
	e.buf = enc.AppendBeginMarker(e.buf)
	e.w = w
	e.level = level
	return e
}

func (e *Event) Str(key, val string) *Event {
	if e == nil {
		return e
	}

	enc := Encoder{}
	e.buf = enc.AppendString(enc.AppendKey(e.buf, key), val)
	return e
}

func (e *Event) Msg(msg string) {
	if e == nil {
		return
	}
	e.msg(msg)
}

func (e *Event) msg(msg string) {
	enc := Encoder{}
	if msg != "" {
		e.buf = enc.AppendString(enc.AppendKey(e.buf, "message"), msg)
	}

	err := e.write()
	if err != nil {
		fmt.Errorf("error %v", err)
		return
	}
}

func (e *Event) write() error {
	if e == nil {
		return nil
	}

	enc := Encoder{}
	e.buf = enc.AppendEndMarker(e.buf)
	e.buf = enc.AppendLineBreak(e.buf)

	_, err := e.w.WriteLevel(e.level, e.buf)
	return err
}
