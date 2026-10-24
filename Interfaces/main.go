package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id    string
	name  string
	cargo int
}

func (t *NormalTruck) LoadCargo() error {
	// Simulate a scenario where the truck is not found.
	t.cargo += 1 // Simulate loading cargo
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	// Simulate a scenario where the truck is not found.
	t.cargo -= 1 // Simulate unloading cargo
	return nil
}

type ElectricTruck struct {
	id      string
	name    string
	battery float64
	cargo   int
}

func (t *ElectricTruck) LoadCargo() error {
	// Simulate a scenario where the truck is not found.
	t.battery -= 1.0 // Simulate battery consumption for loading cargo
	t.cargo += 1     // Simulate loading cargo
	return nil
}

func (t *ElectricTruck) UnloadCargo() error {
	// Simulate a scenario where the truck is not found.
	t.battery -= 0.5 // Simulate battery consumption for unloading cargo
	t.cargo -= 1     // Simulate unloading cargo
	return nil
}

func processTruck(truck Truck) error {
	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading Cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading Cargo: %w", err)
	}

	return nil
}

func main() {
	nt := &NormalTruck{id: "1", name: "Normal Truck 1"}
	fmt.Println("Processing Normal Truck:", *nt)
	et := &ElectricTruck{id: "2", name: "Electric Truck 1"}
	fmt.Printf("Processing the Electric Truck: %+v\n", *et)

	if err := processTruck(nt); err != nil {
		log.Fatalf("error processing normal truck: %v", err)
	}

	if err := processTruck(et); err != nil {
		log.Fatalf("error processing electric truck: %v", err)
	}

	fmt.Println("Normal Truck Cargo:", nt.cargo)
	fmt.Println("Electric Truck Cargo:", et.cargo)
	fmt.Println("Electric Truck Battery:", et.battery)
}
