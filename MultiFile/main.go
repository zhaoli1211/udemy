package main

import (
	"fmt"
)

func main() {
	c := Car{}
	c.numberOfDoor = 4
	c.sylinders = 5
	fmt.Println(c)

	t := Truck{4,4,oneTon}
	fmt.Println(t)
}
