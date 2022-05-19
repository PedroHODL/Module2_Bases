package main

import "fmt"

func main() {
	fmt.Println("--------------Ex 1-------------------")
	Ex_1()
	fmt.Println("\n--------------Ex 2-------------------")
	Ex_2()
	fmt.Println("--------------Ex 3-------------------")
	Ex_3()
	fmt.Println("--------------Ex 4-------------------")
	Ex_4()
}

func Ex_1() {
	var palavra string = "Pneumoultramicroscopicossilicovulcanoconiotico"
	fmt.Println("A palavra tem", len(palavra), "letras")

	for key := range palavra {
		fmt.Printf("%c ", palavra[key])
	}
}

func Ex_2() {
	var (
		idade        = 23
		anoAtividade = 0.5
		empregado    = true
	)

	switch {
	case idade > 22 && empregado == true && anoAtividade >= 1:
		fmt.Println("Você pode solicitar um empréstimo")
	case idade <= 22:
		fmt.Println("Você tem menos de 22 anos. Empréstimo Negado!")
	case empregado == false:
		fmt.Println("Você não está empregado. Empréstimo Negado!")
	case anoAtividade < 1:
		fmt.Println("Você não possui tempo suficiente de atividade. Empréstimo Negado!")
	}
}

func Ex_3() {
	var mes int = 7
	meses := map[int]string{1: "janeiro", 2: "fevereiro", 3: "março", 4: "abril", 5: "maio", 6: "junho",
		7: "julho", 8: "agosto", 9: "setembro", 10: "outubro", 11: "novembro", 12: "dezembro"}
	fmt.Println(meses[mes])
}

func Ex_4() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	cont := 0
	for _, value := range employees {
		if value > 21 {
			cont++
		}
	}
	employees["Frederico"] = 25
	delete(employees, "Pedro")

	fmt.Println(cont, "funcionários tem mais de 21 anos")
	fmt.Println("Novo mapa após alterações:", employees)
}
