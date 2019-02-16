package worker_pool2

import (
	"testing"
)

func TestRunningWorker(t *testing.T) {

	w := &WorkerPool{
		job:    make(chan func() string, 100),
		result: make(chan string, 100),
	}

	w.Start(4)

	for i := 0; i < 100; i++ {
		w.job <- func() string {
			return "가즈아"
		}
	}

	close(w.job)

	for i := 0; i < 100; i++ {
		<-w.result
	}
}
