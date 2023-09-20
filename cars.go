package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	// Short table with 5 columns
	shortTable := table.NewWriter()
	shortTable.SetOutputMirror(nil)
	shortTable.AppendHeader(table.Row{"#", "Car Name", "Make", "Year", "Price"})
	shortTable.AppendRows([]table.Row{
		{1, "Toyota Corolla", "Toyota", 2022, "$24,000"},
		{2, "Honda Civic", "Honda", 2022, "$22,500"},
		{3, "Ford Mustang", "Ford", 2022, "$45,000"},
		{4, "Chevrolet Camaro", "Chevrolet", 2022, "$42,000"},
	})
	fmt.Println("Short Table:")
	fmt.Println(shortTable.Render())

	// Table with 20 columns
	twentyColumnTable := table.NewWriter()
	twentyColumnTable.SetOutputMirror(nil)
	headerRow := table.Row{"#", "Car Name", "Make", "Year", "Price", "Color", "Engine", "Fuel Type", "Mileage", "Transmission",
		"Seats", "Length", "Width", "Height", "Weight", "Top Speed", "0-60 mph", "Horsepower", "Torque", "Drive Type"}
	twentyColumnTable.AppendHeader(headerRow)
	dataRows := []table.Row{
		{1, "Toyota Corolla", "Toyota", 2022, "$24,000", "Silver", "2.0L Inline-4", "Gasoline", "30 mpg", "Automatic",
			"5", "182 in", "70 in", "57 in", "3000 lbs", "120 mph", "8.0 seconds", "169 hp", "151 lb-ft", "Front-Wheel Drive"},
		{2, "Honda Civic", "Honda", 2022, "$22,500", "Blue", "1.5L Inline-4", "Gasoline", "32 mpg", "CVT",
			"5", "177 in", "69 in", "56 in", "2900 lbs", "125 mph", "7.5 seconds", "180 hp", "163 lb-ft", "Front-Wheel Drive"},
		// Add more data rows here...
	}
	twentyColumnTable.AppendRows(dataRows)

	fmt.Println("\n20 Column Table:")
	fmt.Println(twentyColumnTable.Render())
}
use
