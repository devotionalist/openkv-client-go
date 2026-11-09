package kv

import (
	"testing"
)

func TestCasErrorMessage(t *testing.T) {
	err := &CasError{Current: 42}
	if got := err.Error(); got != "cas_failed current_revision=42" {
		t.Fatalf("got %q", got)
	}
}
