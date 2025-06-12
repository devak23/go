package main

// This program just shows how to define an interface type and use it.
type vehicle interface {
	getSpeed() float64
	getDistanceTravelled() float64
}

type car struct {
	brand string
}

func (Car car) getSpeed() float64 {
	return 78.64
}

func (Car car) getDistanceTravelled() float64 {
	return 4000
}

func InterfaceTypesDemo() {
	var c vehicle = car{"Syros"}
	Println("Speed of the vehicle: ", c.getSpeed())
	Println("Distance travelled by the vehicle: ", c.getDistanceTravelled())
}
