package Ex2

import (
	"errors"
	"fmt"
)

func Ex2() {
	salario := 20000

	fmt.Println("Salario:", salario)
	val, err := VerificarImposto(salario)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}

func VerificarImposto(sal int) (string, error) {
	if sal < 15000 {
		return "", errors.New("erro: O salário digitado não está dentro do valor mínimo")
	}
	return "Necessário pagamento de imposto", nil
}
