package main

import "io/ioutil"

func main(){
	input, err := ioutil.ReadFile("./inputtext.txt")
	if err != nil {
		panic(err)
	}
	string1 := string(input)

	var floor int

	for index, value := range string1{
		if string(value) == "("{
			floor++
		}else {
			floor--
		}
		if floor < 0{
			println(index + 1)
		}
	}
}
