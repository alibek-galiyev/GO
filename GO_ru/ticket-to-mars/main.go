package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	headerFmt  = "%-20s %-6s %-10s %-15s"
	rowFmt     = "%-20s %4d   %-12s %16d\n"
	numRows    = 10
	minDays    = 21
	daysRange  = 30
	minPrice   = 30
	priceRange = 100
	lines      = "========================================================="
)

func main() {
	var flightCompanies []string = []string{
		"Virgin Galactic",
		"SpaceX",
		"Space Adventures",
		"AG Corp.",
	}
	var header string = fmt.Sprintf(
		headerFmt,
		"Spaceline", "Days", "Trip type", "Price (in mlns. $)",
	)
	var tripType []string = []string{"Round-trip", "One-way"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println(header)
	fmt.Println(lines)

	for range numRows {
		fmt.Printf(
			rowFmt,
			randomChoice(r, flightCompanies),
			randomIntRange(r, minDays, daysRange),
			randomChoice(r, tripType),
			randomIntRange(r, minPrice, priceRange),
		)
	}

}

func randomIntRange(r *rand.Rand, min, span int) int {
	return r.Intn(span) + min
}

func randomChoice(r *rand.Rand, items []string) string {
	return items[r.Intn(len(items))]
}
