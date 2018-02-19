package main

import (
	"fmt"
	"os"
	"strconv"
)

var in [3]int64
var a, b, c int64

func main() {
	checkInput(os.Args, false)
	findTriangleName(a, b, c)
}

func checkInput(input []string, testing bool) {
	if !testing {
		input = input[1:]
	}
	if len(input) > 3 {
		panic("Too much input. You can only use 3 integers as input!")
	} else if len(input) < 3 {
		fmt.Println("Should panic!")
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
	a = in[0]
	b = in[1]
	c = in[2]
}

func findTriangleName(a, b, c int64) (res string) {
	checkValidTriangle(a, b, c)

	if a == b && b == c {
		fmt.Println("Equilateral triangle")
		return "Equilateral triangle"
	} else if a == b || a == c || b == c {
		fmt.Println("Isosceles triangle")
		return "Isosceles triangle"
	} else {
		fmt.Println("Scalene triangle")
		return "Scalene triangle"
	}
}

func checkValidTriangle(a int64, b int64, c int64) (valid bool) {
	if a+b > c && a+c > b && b+c > a {
		return true
	}
	return false
}
