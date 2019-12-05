package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(calc(input))
}

func calc(input int) int {
	return (input / 3) - 2
}

func calc2(input, total int) (int, int) {
	req := calc(input)
	if calc(req) <= 0 {
		return req, total + req
	}
	return calc2(req, total+req)
}
