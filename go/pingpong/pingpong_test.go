package pingpong

import (
	"sort"
	"testing"
)

func TestRunPingPongLength(t *testing.T) {
	got := RunPingPong(5)
	wantLen := 10

	if len(got) != wantLen {
		t.Fatalf("expected %d messages, got %d", wantLen, len(got))
	}
}

func TestRunPingPongContent(t *testing.T) {
	got := RunPingPong(3)

	// count occurrences
	pingCount := 0
	pongCount := 0

	for _, v := range got {
		if v == "ping" {
			pingCount++
		}
		if v == "pong" {
			pongCount++
		}
	}

	if pingCount != 3 {
		t.Errorf("expected 3 'ping', got %d", pingCount)
	}

	if pongCount != 3 {
		t.Errorf("expected 3 'pong', got %d", pongCount)
	}
}

func TestRunPingPongNoNil(t *testing.T) {
	got := RunPingPong(2)
	for _, v := range got {
		if v == "" {
			t.Fatalf("expected no empty strings, got %+v", got)
		}
	}
}

func TestRunPingPongDeterministicCount(t *testing.T) {
	// order is irrelevant, but total count must match
	N := 8
	got := RunPingPong(N)

	if len(got) != 2*N {
		t.Fatalf("expected %d messages total, got %d", 2*N, len(got))
	}

	// Sorting doesn't affect correctness but helps debugging
	sort.Strings(got)
}
