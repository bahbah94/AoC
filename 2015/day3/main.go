package main

import (
	"io/ioutil"
	"strings"
	"fmt"
)

type coordinate struct{
	x, y int
}

func main(){
		input, err := ioutil.ReadFile("./inputtext.txt")
		if err != nil {
			panic(err)
		}
		home := make(map[coordinate]bool)
		c := coordinate{0,0}

		home[c] = true

		for _, value := range strings.Split(string(input), ""){
			switch value{
			case ">":
				c = coordinate{c.x + 1, c.y}
			case "<":
				c = coordinate{c.x - 1, c.y}
			case "^":
				c = coordinate{c.x, c.y + 1}
			case "v":
				c = coordinate{c.x, c.y - 1}
			}
			home[c] = true

		}
		fmt.Println(len(home))
}
