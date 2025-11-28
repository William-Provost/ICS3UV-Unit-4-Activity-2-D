// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-28
// Fileoverview: This program finds the threshold below which the
// antibiotic is considered expired

package main

import "fmt"

func main() {
	// constants & variables
	const EFFECTIVENESS_THRESHOLD float64 = 50.0
	const DECREASE_RATE float64 = 0.04 // percentage decrease per month
	var effectiveness float64 = 100.0
	var months int = 0

	fmt.Println("Calculating antibiotic effectiveness over time:")
	fmt.Println()
	for effectiveness >= EFFECTIVENESS_THRESHOLD {
		fmt.Printf("Month %d: Effectiveness = %.2f%%\n", months, effectiveness)
		effectiveness = effectiveness - (effectiveness * DECREASE_RATE)
		months++
	}

	// output result
	fmt.Printf("\nThe antibiotic is considered expired after %d months when effectiveness drops below %.2f%%.\n",
		months, EFFECTIVENESS_THRESHOLD)

	fmt.Println("\nDone.")
}
