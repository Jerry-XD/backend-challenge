package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	var arrInput = [][]int{
		// {59},
		// {73, 41},
		// {52, 40, 53},
		// {26, 53, 6, 34},
		// {10, 51, 87, 86, 81},
	}

	data, err := os.ReadFile("hard.json")
	if err != nil {
		log.Panic("Error read file :", err)
	}

	err = json.Unmarshal(data, &arrInput)
	if err != nil {
		log.Panic("Error unmarshal :", err)
	}

	if len(arrInput) == 0 {
		fmt.Println(0)
		return
	}

	for i := len(arrInput) - 2; i >= 0; i-- {
		// log.Println(arrInput[i])
		for j := 0; j <= i; j++ {
			// log.Println(arrInput[i][j])
			// log.Println(arrInput[i+1][j])
			// log.Println(arrInput[i+1][j+1])

			var left = arrInput[i+1][j]
			var right = arrInput[i+1][j+1]

			if left >= right {
				arrInput[i][j] += left
			} else {
				arrInput[i][j] += right
			}
			// log.Println(arrInput[i][j])
		}
	}
	fmt.Println(arrInput[0][0])
}
