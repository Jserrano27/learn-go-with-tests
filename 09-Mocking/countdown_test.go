package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {

	t.Run("prints 3 to Go! with sleeps", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer, &CountdownOperationsSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("Got: %q, but expected was: %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		writeSleepSpy := &CountdownOperationsSpy{}

		Countdown(writeSleepSpy, writeSleepSpy)

		got := writeSleepSpy.Calls
		want := []string{
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, but expected was: %v", got, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 4 * time.Second
	spyTime := SpyTime{}

	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("Got: %v, but expected was: %v", spyTime.durationSlept, sleepTime)
	}
}
