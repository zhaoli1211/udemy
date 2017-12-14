package CarStuff

type Car struct {
	NumberOfDoor int
	Sylinders int
}

func (c *Car) GetDoors() int {
	return c.NumberOfDoor
}