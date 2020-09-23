package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (c *CountdownOperationsSpy) Sleep() {
	c.Calls = append(c.Calls, sleepAction)
}

func (c *CountdownOperationsSpy) Write(w []byte) (n int, err error) {
	c.Calls = append(c.Calls, writeAction)
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

const (
	countdownStart = 3
	finalWord      = "Go!"
	writeAction    = "write"
	sleepAction    = "sleep"
)

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}
	s.Sleep()
	fmt.Fprint(w, finalWord)
}

func main() {
	interval := 1 * time.Second
	sleeper := &ConfigurableSleeper{interval, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
