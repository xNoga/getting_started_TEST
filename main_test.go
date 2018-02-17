package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckInput(t *testing.T) {
	input := []string{"2", "2", "3"}
	fmt.Println("Testing checkInput() function with input 2 2 3")
	res := assert.NotPanics(t, func() { checkInput(input) })
	if !res {
		t.Error("Paniced!")
	}
}

func TestValidTriangleAndName(t *testing.T) {
	fmt.Println("\nTesting checkValidTriangle() and findTriangleNameTest() function with input 2 3 4")

	a := int64(2)
	b := int64(3)
	c := int64(4)

	assert.Equal(t, true, checkValidTriangle(a, b, c), "The triangle was valid!")

	assert.Equal(t, findTriangleNameTest(a, b, c), "Scalene triangle", "The function returned: Scalene triangle")

}

func findTriangleNameTest(a int64, b int64, c int64) (res string) {

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

	return ""
}
