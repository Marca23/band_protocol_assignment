package main

import (
	"fmt"
	"log"
)

// Solve finds the maximum number of chickens that can be protected under a roof of length k.
// n is the number of chickens, k is the length of the roof, and pos is the array of chicken positions.
// The function returns the maximum number of chickens that can be protected.
func Solve(n int, k int, pos []int) int {

	// last keeps the last index of chicken that can still use the roof to protect from the position of the current chicken.
	var last int = 0
	var result int = 0

	// Iterate through the chickens to find the maximum number of chickens that can be protected.
	for i := 0; i < n; i++ {
		// If the distance between the current chicken and the last chicken that can still use the roof is greater than or equal to k,
		// increment the last index until the distance is less than k.
		for pos[i] - pos[last] >= k {
			last++
		}
		// Update the maximum number of chickens that can be protected.
		result = max(result, i - last + 1)	
	}

	return result
}

func main() {
	var n, k, x int
	var pos[]int

	// Read input n, k represent number of chickens and length of roof from the user
	_, err := fmt.Scanln(&n, &k)
	if err != nil {
		log.Fatal(err)
	}
	
	// Read input array of chickens position from the user
	for i := 0; i < n; i++ {
		_, err = fmt.Scanf("%d", &x)
		if err != nil {
			log.Fatal(err)
		}
		pos = append(pos, x)
	}

	// Print the result of the Solve function
	result := Solve(n, k, pos)
	fmt.Println(result)
}