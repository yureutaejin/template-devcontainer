package greeting

import "testing"

func TestMessage(t *testing.T) {
	t.Parallel()

	if got, want := Message("Ada"), "hello, Ada"; got != want {
		t.Fatalf("Message() = %q, want %q", got, want)
	}
}
