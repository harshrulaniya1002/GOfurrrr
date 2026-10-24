package main

import (
	"errors"
	"fmt"
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

}
