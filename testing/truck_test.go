package main

import (
	"testing"
)

// go.mod is necessary for the test to run, as it defines the module path and dependencies.
// It allows the Go toolchain to manage and resolve dependencies correctly during testing.
func TestMainProcess(t *testing.T) {
	t.Run("ProcessTruck", func(t *testing.T) {
		t.Run("Should load and unload a truck cargo", func(t *testing.T) {
			nt := &NormalTruck{id: "1", name: "Normal Truck 1", cargo: 40}
			et := &ElectricTruck{id: "2", name: "Electric Truck 1", battery: 100, cargo: 10}

			if err := processTruck(nt); err != nil {
				t.Fatalf("error processing normal truck: %v", err)
			}

			if err := processTruck(et); err != nil {
				t.Fatalf("error processing electric truck: %v", err)
			}

			//assert the cargo values after processing
			if nt.cargo != 1 {
				t.Errorf("expected normal truck cargo to be 1, got %d", nt.cargo)
			}

			if et.cargo != 0 {
				t.Errorf("expected electric truck cargo to be 0, got %d", et.cargo)
			}

			//assert the battery value after processing electric truck
			if et.battery != -2 {
				t.Errorf("expected electric truck battery to be -2, got %f", et.battery)
			}
		})
	})
}
