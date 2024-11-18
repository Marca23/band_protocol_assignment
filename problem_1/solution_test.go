package main

import "testing"

func TestSolve(t *testing.T) {
    goodBoy := "Good boy"
    badBoy := "Bad boy"
    cases := []struct {
        caseName string
        input, expected string
        err bool 
    }{
        {"Invalid input with characters other than 'R' and 'S'", "ABD", "", true},
        {"All shots revenged", "SRSSRRR", goodBoy, false},
        {"Boss Baby initiates the shot", "RSSRR", badBoy, false},
        {"Boss Baby doesn't revenge the last shot", "SSSRRRRS", badBoy, false},
        {"Unrevenged shots", "SRRSSR", badBoy, false},
        {"All shots revenged with no extra shots", "SSRSRR", goodBoy, false},
        {"Empty input", "", "", true},
        {"Only shots by kids", "SSSS", badBoy, false},
        {"Only revenges by Boss Baby", "RRRR", badBoy, false},
        {"Alternating shots and revenges", "SRSRSRSR", goodBoy, false},
        {"Multiple shots followed by revenges", "SSSSRRRR", goodBoy, false},
    }
    for _, c := range cases {
        result, err := Solve(c.input)
        if err != nil && !c.err {
            t.Errorf("\nCase %q:\ninput: %q\nunexpected error: %v", c.caseName, c.input, err)
        }
        if result != c.expected {
            t.Errorf("\nCase %q:\ninput: %q\nresult %q, expected %q", c.caseName, c.input, result, c.expected)
        }
    }
}