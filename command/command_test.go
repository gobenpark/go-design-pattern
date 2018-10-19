package command

import (
	"fmt"
	"testing"
	"time"
)

func TestCommandQueue_AddCommand(t *testing.T) {
	queue := CommandQueue{}

	queue.AddCommand(CreateCommand("First Message"))
	queue.AddCommand(CreateCommand("Second Message"))
	queue.AddCommand(CreateCommand("Third message"))

	queue.AddCommand(CreateCommand("Fourth message"))
	queue.AddCommand(CreateCommand("Fifth message"))

	t.Run("Command Example2", func(t *testing.T) {
		var timeCommand Command2
		timeCommand = &TimePassed{
			start: time.Now(),
		}

		var helloCommand Command2
		helloCommand = &HelloMessage{}

		time.Sleep(time.Second)

		fmt.Println(timeCommand.Info())
		fmt.Println(helloCommand.Info())

	})

	t.Run("with chain of responsibility ", func(t *testing.T) {
		second := new(Logger)
		first := Logger{NextChain: second}

		command := &TimePassed{start: time.Now()}

		first.Next(command)
	})
}
