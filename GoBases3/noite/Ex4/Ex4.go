package Ex4

import (
	"fmt"
	"math/rand"
	"time"
)

func Ex4() {
	var1 := rand.Perm(100)
	var2 := rand.Perm(1000)
	var3 := rand.Perm(10000)

	var funcoes []func([]int, chan<- string)
	funcoes = append(funcoes, SelectionSort, InsertionSort, BubbleSort)

	ordenarLista(var1, funcoes)
	ordenarLista(var2, funcoes)
	ordenarLista(var3, funcoes)
}

func ordenarLista(array []int, funcoes []func([]int, chan<- string)) {
	c := make(chan string)

	for _, Sort := range funcoes {
		go Sort(array, c)
	}

	for i := 0; i < 3; i++ {
		fmt.Println(<-c)
	}
	fmt.Println()
}

func SelectionSort(array []int, c chan<- string) {
	t := time.Now()

	var tamanho = len(array)
	for i := 0; i < tamanho; i++ {
		var minIdx = i
		for j := i; j < tamanho; j++ {
			if array[j] < array[minIdx] {
				minIdx = j
			}
		}
		array[i], array[minIdx] = array[minIdx], array[i]
	}

	t0 := time.Since(t).Seconds()
	result := fmt.Sprintf("Tempo de execução do Selection Sort no array de %d inteiros: %.6fs",
		tamanho, t0)
	c <- result
}

func InsertionSort(array []int, c chan<- string) {
	t := time.Now()

	var tamanho = len(array)
	for i := 1; i < tamanho; i++ {
		j := i
		for j > 0 {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
			j = j - 1
		}
	}

	t0 := time.Since(t).Seconds()
	result := fmt.Sprintf("Tempo de execução do Insertion Sort no array de %d inteiros: %.6fs",
		tamanho, t0)
	c <- result
}

func BubbleSort(array []int, c chan<- string) {
	t := time.Now()

	var (
		tamanho = len(array)
		sorted  = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < tamanho-1; i++ {
			if array[i] > array[i+1] {
				array[i+1], array[i] = array[i], array[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		tamanho = tamanho - 1
	}

	t0 := time.Since(t).Seconds()
	result := fmt.Sprintf("Tempo de execução do %14s no array de %d inteiros: %.6fs",
		"Bubble Sort", len(array), t0)
	c <- result
}
