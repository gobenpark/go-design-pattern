package object_pool

type Messenger struct{}

func (Messenger) Do() {}

type Pool chan *Messenger

func New(total int) *Pool {
	p := make(Pool, total)
	for i := 0; i < total; i++ {
		p <- new(Messenger)
	}
	return &p
}

func gogo() {
	p := New(4)

	select {
	case obj := <-*p:
		obj.Do()
		*p <- obj
	}
}
