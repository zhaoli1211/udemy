package main

import (
	"fmt"
	cs "github.com/zhaoli1211/udemy/CarStuff"
)

func main() {
	c := cs.Car{}
	c.NumberOfDoor = 4
	c.Sylinders = 5
	fmt.Println(c)
	fmt.Println(c.GetDoors())
	t := Truck{4,4,oneTon}
	fmt.Println(t)
	fmt.Println(t.GetDoors())
}
