package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
)

func main() {
	var r = gin.Default()
	r.GET("/beef/summary", GetBeefSummary)
	_ = r.Run()
}

func GetBeefSummary(c *gin.Context) {
	var url = "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	var method = "GET"
	var client = &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var tokens = tokenize(string(body))
	var sumBeef = map[string]int{}
	for _, v := range tokens {
		sumBeef[v] += 1
	}

	c.JSON(http.StatusOK, gin.H{
		"beef": sumBeef,
	})
}

func tokenize(text string) []string {
	delimiterFunc := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '-'
	}

	tokens := strings.FieldsFunc(strings.ToLower(text), delimiterFunc)
	return tokens
}
