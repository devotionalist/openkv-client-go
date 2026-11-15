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

func TestSplitAddrs(t *testing.T) {
	got := splitAddrs(" 127.0.0.1:19000 , 127.0.0.1:19001 , ")
	if len(got) != 2 || got[0] != "127.0.0.1:19000" || got[1] != "127.0.0.1:19001" {
		t.Fatalf("got %#v", got)
	}
	if splitAddrs("") != nil && len(splitAddrs("")) != 0 {
		t.Fatalf("empty should be empty")
	}
}
