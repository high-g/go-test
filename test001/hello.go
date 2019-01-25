package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var a, b = 2, 3

func main() {
	sayHello()
	sayNumber()
	sayArray()
	sayTime()
}

func sayHello() {
	fmt.Println("Hello World!")
	// global
	fmt.Println(a, b)
}

func sayNumber() {
	a, b := 0, 1
	// local
	fmt.Println(a, b)
	fmt.Println("rand number", rand.Intn(100))
	fmt.Printf("now have you %g problems.\n", math.Sqrt(7))
	fmt.Println("Pi is", math.Pi)
	fmt.Println("add is", add(rand.Intn(100), rand.Intn(100)))
}

func sayArray() {
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		if i < 3 {
			fmt.Printf("p[%d] == %d\n", i, arr[i])
		} else {
			fmt.Println("Yhea")
		}
	}
}

func sayTime() {
	fmt.Println("This time is", time.Now())
}

func add(x, y int) int {
	return x + y
}
