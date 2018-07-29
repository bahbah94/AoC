package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./inputtext.txt")
	if err != nil {
		panic(err)
	}
	string := string(input)

	up := strings.Count(string, "(")
	down := strings.Count(string, ")")

	floor := up - down

	fmt.Println(floor)
}
