package memento

import "fmt"

//Care Taker: 메멘토를 저장하고 디비에 저장할건지 지정된 수를 초과해서 저장하지 않을 건지에 대한 로직을 갖는다
type careTaker struct {
	mementoList []memento
}

func (c *careTaker) Add(m memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *careTaker) Memento(i int) (memento, error) {

	if len(c.mementoList) < i || i < 0 {
		return memento{}, fmt.Errorf("Index net found\n")
	}
	return c.mementoList[i], nil
}
