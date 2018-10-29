package memento

//Originator: Memento를 만들고 현재 활성 state를 저장
type originator struct {
	state State
}

func (o *originator) NewMemento() memento {
	return memento{o.state}
}

func (o *originator) ExtractAndStoreState(m memento) {
	o.state = m.state
}
