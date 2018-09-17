package abstract_factory

type FamilyCar struct{}

func (*FamilyCar) NumWheels() int {
	return 4
}

func (*FamilyCar) NumSeats() int {
	return 5
}

func (*FamilyCar) NumDoors() int {
	return 5
}
