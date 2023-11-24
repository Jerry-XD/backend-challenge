package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Print("enter code : ")
	code, _ := reader.ReadString('\n')

	code = strings.ReplaceAll(strings.ToUpper(code), " ", "")
	fmt.Print("code = " + code)

	var num = make([]int, len(code))
	var codeArr = strings.Split(code, "")

	var result = solve(codeArr, num)

	log.Println(result)
}

func solve(codeArr []string, num []int) string {
	for i, v := range codeArr {
		// log.Println(i, v)
		if v == "L" {
			// log.Println("case L")
			if i > 0 {
				if num[i] <= num[i-1] { // if current < left
					num[i] += num[i+1] + 1  // add current
					num[i-1] += num[i]      // add left
					if num[i-1] == num[i] { // if left == current
						num[i-1] += 1 // add left
					}

				}

			}
		} else if v == "R" {
			// log.Println("case R")
			num[i+1] += num[i] + 1

		} else if v == "=" {
			// log.Println("case =")
			num[i+1] = num[i]
		}
		// log.Println(num)
	}
	return arrayToString(num)
}

func arrayToString(a []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), ""), "[]")
}
