package main

import "fmt"

func main() {
	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					// select로 감싼이유는 valStream에 대한 수신하는 부분이 없는경우 block 이 되기때문
					select {
					case valStream <- v:
					case <-done:
					}
				}
		}()
		return valStream
	}

	done := make(chan interface{})
	myChan := make(chan interface{})

	for val := range orDone(done, myChan) {
		fmt.Println(val)
	}
}
