package main

import "testing"

func TestSolve(t *testing.T) {
    cases := []struct {
        caseName string
        n, k int
        pos []int
        expected int
    }{
        {"Case 1", 5, 5, []int{2, 5, 10, 12, 15}, 2},
        {"Case 2", 6, 10, []int{1, 11, 30, 34, 35, 37}, 4},
        {"Case 3", 4, 3, []int{1, 2, 3, 4}, 3},
        {"Case 4", 3, 7, []int{3, 7, 14}, 2},
        {"Case 5", 7, 25, []int{5, 10, 15, 20, 25, 30, 35}, 5},
        {"Case 6", 8, 25, []int{2, 5, 10, 15, 20, 25, 30, 35}, 6},
        {"Case 7", 5, 1, []int{1, 2, 3, 4, 5}, 1},
        {"Case 8", 10, 123456789, []int{500000, 1000000, 1500000, 2000000, 2500000, 3000000, 3500000, 4000000, 4500000, 5000000}, 10},    
    }
    for _, c := range cases {
        result := Solve(c.n, c.k, c.pos)
        if result != c.expected {
            t.Errorf("\n%q:\ninput: n=%d, k=%d, pos=%d\nresult %d, expected %d", c.caseName, c.n, c.k, c.pos, result, c.expected)
        }
    }
}