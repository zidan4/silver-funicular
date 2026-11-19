package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	age int = 30
	name string = "John Doe"
)



func main( ) {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) // Generate a random number between 0 and 99
	fmt.Println("Random Number:", randomNumber)

	var mat complex128 = complex(float64(randomNumber), float64(randomNumber))
	fmt.Println("Complex Number:", mat)
	fmt.Println(imag(mat))
}