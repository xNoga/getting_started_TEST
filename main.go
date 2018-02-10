package main

import (
	"fmt"
	"os"
	"strconv"
)

var in [3]int64

func main() {
	checkInput(os.Args)
	findTriangleName()
}

func checkInput(input []string) {

	if len(input) > 4 {
		panic("Too much input. You can only use 3 integers as input!")
	}

	if len(input) < 4 {
		panic("Not enough input. You need to input 3 integers as input!")
	}

	for i, v := range input[1:] {
		val, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			fmt.Println("Invalid input. The input " + v + " was not an integer")
			panic(err)
		} else {
			in[i] = val
		}
	}
}

func findTriangleName() {
	var a = in[0]
	var b = in[1]
	var c = in[2]

	if a == b && b == c {
		fmt.Println("Equilateral triangle")
	} else if a == b || a == c || b == c {
		fmt.Println("Isosceles triangle")
	} else {
		fmt.Println("Scalene triangle")
	}
}
