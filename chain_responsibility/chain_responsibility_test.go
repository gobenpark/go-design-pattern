package chain_responsibility

import (
	"fmt"
	"strings"
	"testing"
)

func TestCreateDefaultChain(t *testing.T) {
	myWriter := myTestWriter{}

	writerLogger := WriterLogger{Writer: &myWriter}
	second := SecondLogger{NextChain: &writerLogger}
	chain := FirstLogger{NextChain: &second}

	t.Run("3 loggers, 2 of them writes to console, second only if it founds "+
		"the word 'hello', third writes to some variable if second found 'hello'", func(t *testing.T) {
		chain.Next("Message that breaks the chain\n")

		if myWriter.receivedMessage != nil {
			t.Fatal("Last link should not receive any message")
		}

		chain.Next("Hello\n")

		if myWriter.receivedMessage == nil || !strings.Contains(*myWriter.receivedMessage, "Hello") {
			t.Fatal("Last link didn't reveived expected message")
		}
	})

	t.Run("2 loggers, second uses the closure implementation", func(t *testing.T) {
		myWriter := myTestWriter{}

		closureLogger := ClosureChain{
			Closure: func(s string) {
				fmt.Printf("My closure logger! Message: %s\n", s)
				myWriter.receivedMessage = &s
			},
		}
		writerLogger.NextChain = &closureLogger

		chain.Next("Hello closure logger")
		if *myWriter.receivedMessage != "Hello closure logger" {
			t.Fatal("Expected message wasn't reveive in myWriter")
		}

	})
}
