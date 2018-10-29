package memento

//객체의 상태 정보를 저장 및 복원하는 패턴

//Memento: state를 저장하는 유형
type memento struct {
	state State
}
