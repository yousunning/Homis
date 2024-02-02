package main

import "fmt"

// 추상팩토리 패턴
type Vehicle interface {
	GetVehicleType() string
}

type Car struct {
}

func (c *Car) GetVehicleType() string {
	return "Car"
}

type Truck struct {
}

func (t *Truck) GetVehicleType() string {
	return "Truck"
}

type VehicleFactory interface {
	NewVehicle() Vehicle
}

type CarFactory struct {
}

func (cf *CarFactory) NewVehicle() Vehicle {
	return &Car{}
}

type TruckFactory struct {
}

func (tf *TruckFactory) NewVehicle() Vehicle {
	return &Truck{}
}

func GetFactory(factoryType string) VehicleFactory {
	if factoryType == "car" {
		return &CarFactory{}
	} else if factoryType == "truck" {
		return &TruckFactory{}
	}

	return nil
}

func main() {
	carFactory := GetFactory("car")
	truckFactory := GetFactory("truck")

	car := carFactory.NewVehicle()
	truck := truckFactory.NewVehicle()

	fmt.Println(car.GetVehicleType())
	fmt.Println(truck.GetVehicleType())
}
