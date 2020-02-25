package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

// CountdownOperationsSpy implements both io.Writer and countdown.SpySleeper.
type CountdownOperationsSpy struct {
	Calls []string
}

func (c *CountdownOperationsSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {

	t.Run("test for output", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer, &CountdownOperationsSpy{})

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		got := spySleepPrinter.Calls
		want := []string{
			sleep, write,
			sleep, write,
			sleep, write,
			sleep, write,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted calls %v got %v", want, got)
		}
	})

}
