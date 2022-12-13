package greeter

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Sophie")

	got := buffer.String()
	want := "Hello, Sophie"

	if got != want {
		t.Errorf("Got %q but wanted %q", got, want)
	}
}
