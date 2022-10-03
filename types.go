package bulog

type Encoder struct{}

func (Encoder) AppendBeginMarker(dst []byte) []byte {
	return append(dst, '{')
}

func (Encoder) AppendEndMarker(dst []byte) []byte {
	return append(dst, '}')
}

func (Encoder) AppendLineBreak(dst []byte) []byte {
	return append(dst, '\n')
}

func (e Encoder) AppendString(dst []byte, s string) []byte {
	dst = append(dst, '"')

	dst = append(dst, s...)

	return append(dst, '"')
}

func (e Encoder) AppendKey(dst []byte, key string) []byte {
	if dst[len(dst)-1] != '{' {
		dst = append(dst, ',')
	}

	return append(e.AppendString(dst, key), ':')
}
