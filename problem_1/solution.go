package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// Solve determines if Boss Baby has revenged all shots from kids without initiating any, based on the input string of 'R' and 'S' where:
// 'R' represents a shot revenged by Boss Baby.
// 'S' represents a shot by kids.
// returns "Good boy" if all shots are revenged and no shots are initiated by Boss Baby.
// returns "Bad boy" if there are unrevenged shots or if Boss Baby initiates a shot.
func Solve(s string) (string, error) {
	
	// Validate the input string to ensure it only contains 'R' and 'S'
	s = strings.TrimSpace(s)
	if s == "" {
		return "", errors.New("invalid input: empty string")
	}
	for _, char := range s {
		if char != 'R' && char != 'S' {
			
			return "", errors.New("invalid input: string contains characters other than 'R' and 'S'")
		}
	}

	unrevengedShots := 0
	
	// If the first character is 'R', return "Bad boy" immediately
	if s[0] == 'R' {
		return "Bad boy", nil
	}
	
	// Iterate through the string to count unrevenged shots
	for _, char := range s {
		if char == 'S' {
			unrevengedShots++
		} else if char == 'R' {
			unrevengedShots--
		}
		// Ensure unrevengedShots does not go below 0
		if unrevengedShots < 0 {
			unrevengedShots = 0
		}
	}
	
	// If there are any unrevenged shots left, return "Bad boy"
	if unrevengedShots > 0 {
		return "Bad boy", nil
	}
	
	// Otherwise, return "Good boy"
	return "Good boy", nil
}

func main() {
	var s string
	
	// Read input string s from the user
	_, err := fmt.Scanln(&s)
	if err != nil {
		log.Fatal(err)
	}

	// Print the result of the Solve function
	result, err := Solve(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}