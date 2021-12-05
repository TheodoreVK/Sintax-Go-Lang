package main

import "fmt"

//kode Program dengan Looping
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

//kode Program dengan Rekursif
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
func rekursif() {
	fmt.Println(factorialLoop(5))
	fmt.Println(factorialRecursive(3))

}
