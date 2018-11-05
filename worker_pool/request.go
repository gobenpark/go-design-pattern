package worker_pool

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"sync"
)

type RequestHandler func(interface{})

type Request struct {
	Data    interface{}
	Handler RequestHandler
}

func NewStringRequest(s string, wg *sync.WaitGroup) Request {
	myRequest := Request{
		Data: s,
		Handler: func(i interface{}) {
			defer wg.Done()
			s, ok := i.(string)
			if !ok {
				log.Fatal("Invalid casting to string")
			}
			fmt.Println(s)
		},
	}
	return myRequest
}
