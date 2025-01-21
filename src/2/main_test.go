package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	var inputCaseA = strings.Split("LLRR=", "")
	var inputCaseB = strings.Split("==RLL", "")
	var inputCaseC = strings.Split("=LLRR", "")
	var inputCaseD = strings.Split("RRL=R", "")

	var expectCaseA = "210122"
	var expectCaseB = "000210"
	var expectCaseC = "221012"
	var expectCaseD = "012001"

	var resultA = solve(inputCaseA, make([]int, 6))
	var resultB = solve(inputCaseB, make([]int, 6))
	var resultC = solve(inputCaseC, make([]int, 6))
	var resultD = solve(inputCaseD, make([]int, 6))

	assert.Equal(t, expectCaseA, resultA)
	assert.Equal(t, expectCaseB, resultB)
	assert.Equal(t, expectCaseC, resultC)
	assert.Equal(t, expectCaseD, resultD)
}
