package greet

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	name := "Chris"
	Greet(&buffer, name)

	got := buffer.String()
	want := "Hello, " + name

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
