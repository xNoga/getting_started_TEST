package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckInput(t *testing.T) {
	input := []string{"2", "2", "3"}
	fmt.Println("Testing checkInput() function with input 2 2 3")
	res := assert.NotPanics(t, func() { checkInput(input, true) }, "The code should not panic")
	if !res {
		t.Error("Paniced!")
	}
}

func TestCheckInputFailure(t *testing.T) {
	fmt.Println("\nTesting checkInput() function with input 2 2 3 4")
	input := []string{"2", "2", "3", "4"}
	assert.Panics(t, func() { checkInput(input, true) }, "The code should panic")
}

func TestCheckValidTriangle(t *testing.T) {
	fmt.Println("\nTesting checkValidTriangle() function with input 2 2 3")
	a := int64(2)
	b := int64(2)
	c := int64(3)
	assert.Equal(t, true, checkValidTriangle(a, b, c), "Should be valid triangle")
}

func TestCheckValidTriangleFailure(t *testing.T) {
	fmt.Println("\nTesting checkValidTriangle() function with input 2 2 3")
	a := int64(1)
	b := int64(2)
	c := int64(3)
	assert.Equal(t, false, checkValidTriangle(a, b, c), "Should not be valid triangle")
}

func TestValidTriangleAndName(t *testing.T) {
	fmt.Println("\nTesting findTriangleNameTest() function with input 2 3 4")
	a := int64(2)
	b := int64(3)
	c := int64(4)
	assert.Equal(t, findTriangleName(a, b, c), "Scalene triangle", "The function returned: Scalene triangle")
}
