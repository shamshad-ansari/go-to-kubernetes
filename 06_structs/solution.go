//go:build ignore

package main

import "fmt"

type Engine struct {
	HorsePower   int
	FuelCapacity float64
}

type Vehicle struct {
	Make    string
	Model   string
	Year    int
	Mileage float64
	MPG     float64
	Engine  Engine
}

func (v Vehicle) FuelRange() float64 {
	return v.Engine.FuelCapacity * v.MPG
}

func (v Vehicle) NeedsService() bool {
	return v.Mileage > 100000
}

func (v Vehicle) String() string {
	return fmt.Sprintf("%d %s %s (%.1f mi, range: %.1f mi)",
		v.Year, v.Make, v.Model, v.Mileage, v.FuelRange())
}

func FleetSummary(vehicles []Vehicle) string {
	if len(vehicles) == 0 {
		return "Fleet: 0 vehicles"
	}

	var totalRange float64
	needService := 0
	for _, v := range vehicles {
		totalRange += v.FuelRange()
		if v.NeedsService() {
			needService++
		}
	}
	avgRange := totalRange / float64(len(vehicles))

	return fmt.Sprintf("Fleet: %d vehicles | Avg range: %.1f mi | Need service: %d",
		len(vehicles), avgRange, needService)
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
