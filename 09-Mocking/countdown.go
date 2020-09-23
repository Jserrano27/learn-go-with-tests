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

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (d DefaultSleeper) Sleep() {
	time.Sleep(countdownInterval)
}

const (
	countdownStart    = 3
	countdownInterval = 1 * time.Second
	finalWord         = "Go!"
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
	s := &DefaultSleeper{}
	Countdown(os.Stdout, s)
}
