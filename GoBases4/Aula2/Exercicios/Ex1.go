package Exercicios

import (
	"fmt"
	"os"
)

func Ex1() {
	file, err := os.ReadFile("costumers.txt")
	if err != nil {
		panic("o arquivo indicado não foi encontrado ou está danificado")
	}
	fmt.Println(file)
}
