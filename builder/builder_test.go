package builder

import "testing"

func TestBuilderPattern(t *testing.T) {

	// 싱글톤으로 manufactor를 만들면 글로벌하게 사용이 가능하다.
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}

	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	bike := bikeBuilder.GetVehicle()

	if bike.Wheels != 2 {
		t.Errorf("Wheels on a bike must be 2 and they were %d\n", bike.Wheels)
	}

	if bike.Structure != "Motorbike" {
		t.Errorf("Structure on a bike must be 'Motorbike' and was %s\n", bike.Structure)
	}

	if bike.Seats != 2 {
		t.Errorf("Seats on a bike must be 2 and they were %d\n", bike.Seats)
	}

}
