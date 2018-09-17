package builder

type BuilderProcess interface {
	SetWheels() BuilderProcess
	SetSeats() BuilderProcess
	SetStructure() BuilderProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct {
	builder BuilderProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

func (f *ManufacturingDirector) SetBuilder(b BuilderProcess) {
	f.builder = b
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetStructure() BuilderProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

func (c *CarBuilder) SetWheels() BuilderProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuilderProcess {
	c.v.Seats = 5
	return c
}

type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuilderProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuilderProcess {
	b.v.Seats = 2
	return b
}

func (b *BikeBuilder) SetStructure() BuilderProcess {
	b.v.Structure = "Motorbike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}
