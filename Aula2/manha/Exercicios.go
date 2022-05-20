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
	func1 := FuncionariosS{1, 162.0, "C"}
	func2 := FuncionariosS{2, 176.0, "B"}
	func3 := FuncionariosS{3, 172.0, "A"}

	salario(func1)
	salario(func2)
	salario(func3)
}

type FuncionariosS struct {
	Num   int
	Horas float64
	Tipo  string
}

func salario(funcionario FuncionariosS) {
	res := 0.0

	switch funcionario.Tipo {
	case "A":
		res += funcionario.Horas * 3000.0
		if funcionario.Horas > 160.0 {
			res *= 1.5
		}
	case "B":
		res += funcionario.Horas * 1500.0
		if funcionario.Horas > 160.0 {
			res *= 1.2
		}
	case "C":
		res += funcionario.Horas * 1000.0
	}
	fmt.Printf("O funcionário %d do tipo %s receberá um salário de $%.2f\n",
		funcionario.Num, funcionario.Tipo, res)
}
