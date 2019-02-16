package worker_pool2

import "fmt"

func RunningWorker(workerNum int, jobs <-chan func() string, result chan<- string) {
	for job := range jobs {
		fmt.Printf("%d 디스패치\n", workerNum)
		result <- job()
	}
}

type WorkerPool struct {
	job    chan func() string
	result chan string
}

func (w *WorkerPool) Start(count int) {
	for i := 0; i < count; i++ {
		go RunningWorker(i, w.job, w.result)
	}
}
