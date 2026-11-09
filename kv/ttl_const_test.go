// Copyright 2026 devotionalist
package kv

import "testing"

func TestTTLConstants(t *testing.T) {
	if TTLMissing != -2 || TTLNoExpire != -1 {
		t.Fatalf("unexpected TTL codes missing=%d noexpire=%d", TTLMissing, TTLNoExpire)
	}
}
