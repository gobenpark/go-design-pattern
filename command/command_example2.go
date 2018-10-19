package command

import (
	"fmt"
	"time"
)

type Command2 interface {
	Info() string
}

type TimePassed struct {
	start time.Time
}

func (t *TimePassed) Info() string {
	return time.Since(t.start).String()
}

type HelloMessage struct{}

func (h HelloMessage) Info() string {
	return "Hello world!"
}

// with chain of responsibility

type ChainLogger interface {
	Next(Command2)
}

type Logger struct {
	NextChain ChainLogger
}

func (f *Logger) Next(c Command2) {
	time.Sleep(time.Second)

	fmt.Printf("Elapsed time from creation: %s\n", c.Info())

	if f.NextChain != nil {
		f.NextChain.Next(c)
	}
}
