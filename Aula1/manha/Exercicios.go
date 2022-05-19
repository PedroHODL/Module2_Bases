package main

import "fmt"

func main() {
	Ex_1()
	Ex_2()
}

func Ex_1() {
	var nome string = "Pedro"
	var idade int = 20
	fmt.Println("\nMeu nome é", nome, "e tenho", idade, "anos")
}

func Ex_2() {
	var temperatura int = 9
	var umidade int = 71
	var pressao int = 1017
	fmt.Println("\nA temperatura é de ", temperatura, " graus",
		"\na umidade do ar é de ", umidade, "%",
		"\ne a pressão é de ", pressao, "mbar")
}
