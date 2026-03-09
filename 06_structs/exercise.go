package main

import "fmt"

// Engine represents a vehicle's engine.
type Engine struct {
	HorsePower   int
	FuelCapacity float64
}

// Vehicle represents a car in the fleet.
type Vehicle struct {
	Make    string
	Model   string
	Year    int
	Mileage float64
	MPG     float64
	Engine  Engine
}

// FuelRange returns how many miles the vehicle can travel on a full tank.
// Formula: Engine.FuelCapacity * MPG
func (v Vehicle) FuelRange() float64 {
	// TODO: Return Engine.FuelCapacity * MPG.

	return 0 // replace this
}

// NeedsService returns true if the vehicle has more than 100,000 miles.
func (v Vehicle) NeedsService() bool {
	// TODO: Return true if Mileage > 100000.

	return false // replace this
}

// String returns a one-line summary of the vehicle.
// Format: "%d %s %s (%.1f mi, range: %.1f mi)"
// Example: "2020 Honda Civic (125000.0 mi, range: 372.0 mi)"
func (v Vehicle) String() string {
	// TODO: Return a formatted string using fmt.Sprintf.
	// Include Year, Make, Model, Mileage, and FuelRange().

	return "" // replace this
}

// FleetSummary takes a slice of vehicles and returns a summary string.
//
// Format:
//
//	Fleet: X vehicles | Avg range: Y.1f mi | Need service: Z
//
// Where X is total count, Y is average fuel range, Z is count needing service.
func FleetSummary(vehicles []Vehicle) string {
	// TODO: Calculate total fuel range across all vehicles.
	// TODO: Calculate average fuel range.
	// TODO: Count vehicles where NeedsService() is true.
	// TODO: Return the formatted summary string.

	return "" // replace this
}

func main() {
	fleet := []Vehicle{
		{
			Make: "Honda", Model: "Civic", Year: 2020,
			Mileage: 125000, MPG: 30,
			Engine: Engine{HorsePower: 158, FuelCapacity: 12.4},
		},
		{
			Make: "Toyota", Model: "Camry", Year: 2019,
			Mileage: 85000, MPG: 28,
			Engine: Engine{HorsePower: 203, FuelCapacity: 15.8},
		},
		{
			Make: "Ford", Model: "F-150", Year: 2021,
			Mileage: 45000, MPG: 20,
			Engine: Engine{HorsePower: 290, FuelCapacity: 23.0},
		},
		{
			Make: "Tesla", Model: "Model 3", Year: 2022,
			Mileage: 30000, MPG: 132,
			Engine: Engine{HorsePower: 283, FuelCapacity: 0.0},
		},
	}

	fmt.Println("=== Fleet Vehicles ===")
	for _, v := range fleet {
		fmt.Printf("  %s  |  Service needed: %t\n", v.String(), v.NeedsService())
	}

	fmt.Println("\n=== Fleet Summary ===")
	fmt.Println(FleetSummary(fleet))
}
