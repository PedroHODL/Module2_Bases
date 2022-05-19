package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Scanln(&n1, &n2)
	result := mdc_euclides(n1, n2)
	fmt.Println("O mdc de", n1, "e", n2, "é", result)

	primes := []string{"A", "B", "C", "D"}
	primes[0] = "Hello"
	primes[1] = "World"
	primes[2] = "!"
	fmt.Println(primes)
}

func mdc_euclides(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return mdc_euclides(b, a%b)
	}
}
