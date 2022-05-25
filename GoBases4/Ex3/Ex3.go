package Ex3

import (
	"fmt"
)

func Ex3() {
	salario := 12000

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
		err := fmt.Errorf("erro: O mínimo tributável é $15000 e o valor informado é $%d", sal)
		return "", err
	}
	return "Necessário pagamento de imposto", nil
}
