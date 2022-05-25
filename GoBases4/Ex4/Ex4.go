package Ex4

import (
	"errors"
	"fmt"
)

func Ex4() {
	horas := 79.0
	valor := 120.0

	fmt.Printf("Valor do salário antes do imposto: $%.2f \n", horas*valor)
	val, err := CalcularImposto(horas, valor)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("O valor do salário após aplicar o imposto é de: %.2f \n", val)
}

func CalcularImposto(horas, valor float64) (float64, error) {
	if horas < 80 {
		return 0.0, errors.New("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês")
	}

	salario := horas * valor

	if salario > 20000.0 {
		salario *= 0.9
	}

	return salario, nil
}
