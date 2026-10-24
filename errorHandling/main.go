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

type Truck struct {
	id   string
	name string
}

func (t *Truck) LoadCargo() error {
	// Simulate a scenario where the truck is not found.
	return nil
}

func (t *Truck) UnloadCargo() error {
	// Simulate a scenario where the truck is not found.
	return nil
}

func processTruck(truck Truck) error {
	fmt.Printf("Processing the truck with ID %s and name %s\n", truck.id, truck.name)
	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading Cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("Error unloading Cargo: %w", err)
	}

	return nil
}

func main() {
	trucks := []Truck{
		{id: "1", name: "Truck 1"},
		{id: "2", name: "Truck 2"},
		{id: "3", name: "Truck 3"},
	}

	for _, truck := range trucks {
		if err := processTruck(truck); err != nil {
			if errors.Is(err, ErrTruckNotFound) {
				log.Fatal("Truck not found")
			}

			log.Fatalf("error Processing truck: %s", err)
		}
	}
}
