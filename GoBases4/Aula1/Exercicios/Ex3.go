package Exercicios

import (
	"fmt"
)

func Ex3() {
	salario := 12000

	fmt.Println("Salario:", salario)
	val, err := VerificarImposto3(salario)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}

func VerificarImposto3(sal int) (string, error) {
	if sal < 15000 {
		err := fmt.Errorf("erro: O mínimo tributável é $15000 e o valor informado é $%d", sal)
		return "", err
	}
	return "Necessário pagamento de imposto", nil
}
