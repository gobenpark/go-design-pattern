package composite

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

func Swim() {
	fmt.Println("Swimming!")
}

type Animal struct{}

func (r *Animal) Eat() {
	println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}

// 인터페이스를 사용한 컴포지트 패턴
type Swimmer interface {
	Swim()
}
type Trainer interface {
	Train()
}

type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() {
	println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer // 필드명을 넣지 않을경우 go에서는 상속과 같이 보이도록 객체내에 객체를 포함
	Swimmer
}

func main() {
	swimmer := CompositeSwimmerA{
		MySwim: Swim,
	}
	swimmer.MyAthlete.Train()
	swimmer.MySwim()
	Shark{}.Eat()
	Shark{}.Swim()
	CompositeSwimmerB{}.Train()
	CompositeSwimmerB{}.Swim()
}
