package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	var caseA = strings.Split("LLRR=", "")
	var caseB = strings.Split("==RLL", "")
	var caseC = strings.Split("=LLRR", "")

	var expectCaseA = "210122"
	var expectCaseB = "000210"
	var expectCaseC = "221012"

	var resultA = solve(caseA, make([]int, 6))
	var resultB = solve(caseB, make([]int, 6))
	var resultC = solve(caseC, make([]int, 6))

	if resultA != expectCaseA {
		t.Fatalf("case a error")
	}

	if resultB != expectCaseB {
		t.Fatalf("case b error")
	}

	if resultC != expectCaseC {
		t.Fatalf("case c error")
	}

}
