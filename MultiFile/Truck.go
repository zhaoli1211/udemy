package main

type Truck struct {
	numberOfDoor int
	bedSize int
	weightBtTon weight
}

type weight string

const (
	oneTon weight = "One Ton"
	towTon weight = "Two Ton"
	threeTon weight = "Three Ton"
)

func (t *Truck) GetDoors() int {
	return t.numberOfDoor
}