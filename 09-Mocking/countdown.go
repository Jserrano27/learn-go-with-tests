package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countdownStart    = 3
	countdownInterval = 1 * time.Second
	finalWord         = "Go!"
)

func Countdown(w io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		time.Sleep(countdownInterval)
	}

	fmt.Fprint(w, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
