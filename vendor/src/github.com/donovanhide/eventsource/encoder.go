package eventsource

import (
	"fmt"
	"io"
	"strings"
)

var (
	encFields = []struct {
		prefix string
		value  func(Event) string
	}{
		{"id: ", Event.Id},
		{"event: ", Event.Event},
		{"data: ", Event.Data},
	}
)

type encoder struct {
	w io.Writer
}

func newEncoder(w io.Writer) *encoder {
	return &encoder{w: w}
}

func (enc *encoder) Encode(ev Event) (err error) {
	for _, field := range encFields {
		prefix, value := field.prefix, field.value(ev)
		if len(value) == 0 {
			continue
		}
		value = strings.Replace(value, "\n", "\n"+prefix, -1)
		if _, err = io.WriteString(enc.w, prefix+value+"\n"); err != nil {
			err = fmt.Errorf("Eventsource: Encode: %s", err)
			return
		}
	}
	if _, err = io.WriteString(enc.w, "\n"); err != nil {
		err = fmt.Errorf("Eventsource: Encode: %s", err)
	}
	return
}
