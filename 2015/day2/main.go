package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
)

func main(){
	input, err := ioutil.ReadFile("./inputtext.txt")
	if err != nil {
		panic(err)
	}
	presents := strings.Split(string(input), "\n")

	total := 0

	for _, present := range presents{
		sides := []int{}

		for _, sideString := range strings.Split(present, "x"){
			side, _ := strconv.Atoi(sideString)
			sides = append(sides, side)
		}
		sort.Ints(sides)

		total += 3 * sides[0] * sides[1]
		total += 2 * sides[1] * sides[2]
		total += 2 * sides[0] * sides[2]
	}
	println(total)
}
