package main

import "fmt"

func main() {
	factorial(5)
}

func factorial(x int) {
	var z int = 1
	if x > 0 {
		for i := 1; i <= x; i++ {
			z = z * i
		}
		fmt.Println(x, "Factorial = ", z)
	} else {
		fmt.Println("0'dan küçük olamaz")
	}

}
