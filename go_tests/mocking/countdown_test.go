package main

import (
	"bytes"
	"testing"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	actual := buffer.String()
	expected := `3
2
1
Go!`

	if actual != expected {
		t.Errorf("expected %q got %q", expected, actual)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper. Wants 3 got %d", spySleeper.Calls)
	}

}
