package main

import (
	"fmt"
	"math/rand"
)

// Game Variables
var balls int = 120
var overs int = 0
var wickets int = 0
var runs int = 0

func main() {
	fmt.Println("Golang - cricket simulation")

	w := fmt.Sprintf("wickets %d", wickets)
	r := fmt.Sprintf("runs %d", runs)
	b := fmt.Sprintf("balls %d", balls)
	// o := fmt.Sprintf("overs %d", overs)

	fmt.Println(w)
	fmt.Println(b)
	fmt.Println(r)
	// fmt.Println(o)

	for i := 0; i < 120; i++ {
		s := random()
		runs += s
		print := fmt.Sprintf("Ball %d , Result: %d", i, s)
		fmt.Println(print)
		if i%6 == 0 {
			fmt.Println("End of over")
		}
		balls--

	}

	fmt.Println("Final Runs: ")
	fmt.Println(runs)

}

// Outcome of bowl
func random() int {
	var outcome [5]int
	outcome[0] = 1
	outcome[1] = 2
	outcome[2] = 4
	outcome[3] = 6
	outcome[4] = 0

	min := 0
	max := 5

	return outcome[rand.Intn(max-min)+min]
}
