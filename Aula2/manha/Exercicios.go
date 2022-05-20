package main

import (
	"errors"
	"fmt"
)

func main() {
	Ex_1()
	fmt.Println("--------------Ex 2-------------------")
	Ex_2()
	fmt.Println("--------------Ex 3-------------------")
	Ex_3()
	fmt.Println("--------------Ex 4-------------------")
}

func Ex_1() {
	salarios := map[string]int{"Func1": 50000, "Func2": 150000}

	for _, value := range salarios {
		switch {
		case value == 50000:
			fmt.Println("O imposto é de 17%")
		case value == 150000:
			fmt.Println("O imposto é de 27%")
		default:
			fmt.Println("Erro")
		}
	}
}

func Ex_2() {
	aluno1, _ := media(7, 7, 7, 7)
	aluno2, _ := media(5, 6, 7, 8, 9)
	_, err3 := media(-2, 3, 10)

	fmt.Println("média do aluno 1:", aluno1)
	fmt.Println("média do aluno 2:", aluno2)
	fmt.Println("erro no aluno3:", err3)
}

func media(notas ...int) (int, error) {
	res := 0
	for _, nota := range notas {
		if nota >= 0 {
			res += nota
		} else {
			err := errors.New("Nota negativa encontrada")
			return 0, err
		}
	}
	return res / len(notas), nil
}

func Ex_3() {
	funcionarios := []map[string]interface{}{
		{
			"num":   1,
			"horas": 162.0,
			"tipo":  "C",
		},
		{
			"num":   2,
			"horas": 176.0,
			"tipo":  "B",
		},
		{
			"num":   3,
			"horas": 172.0,
			"tipo":  "A",
		},
	}

	for _, funcionario := range funcionarios {
		res := 0.0

		switch funcionario["tipo"] {
		case "A":
			res += funcionario["horas"].(float64) * 3000.0
			if funcionario["horas"].(float64) > 160.0 {
				res *= 1.5
			}
		case "B":
			res += funcionario["horas"].(float64) * 1500.0
			if funcionario["horas"].(float64) > 160.0 {
				res *= 1.2
			}
		case "C":
			res += funcionario["horas"].(float64) * 1000.0
		}

		fmt.Printf("O funcionário %d do tipo %s receberá um salário de $%.2f\n",
			funcionario["num"], funcionario["tipo"], res)
	}
}
