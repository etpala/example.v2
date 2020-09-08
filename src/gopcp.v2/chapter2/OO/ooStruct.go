package main

import "fmt"

//定义一个车
type Car struct {
	weight int
	name   string
}

//让车可以跑的方法是通过
func (self *Car) Run() {
	fmt.Println(self.name, "is running")
}

type BikeCar struct {
	Car
	lunzi int
}

type GiantBike struct {
	BikeCar
	material string
}

func main() {
	var a GiantBike
	a.weight = 100
	a.name = "Giant bike"
	a.lunzi = 2
	a.material = "carbon"

	fmt.Println(a)
	a.Run()

	var b BikeCar
	b.weight = 100
	b.name = "sample Bike"

	fmt.Println(b)
	b.Run()
	fmt.Println("----")
	var c Car
	c = b.Car
	c.Run()
	c = a.Car
	c.Run()
	a.Car = b.Car
	a.Run()

}
