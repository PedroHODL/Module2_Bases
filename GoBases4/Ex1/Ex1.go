package Ex1

import "fmt"

func Ex1() {
	salario := 10000

	fmt.Println("Salario:", salario)
	val, err := VerificarImposto(salario)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}

type Error struct {
	msg string
}

func (e Error) Error() string {
	return fmt.Sprintf("erro: %s", e.msg)
}

func VerificarImposto(sal int) (string, error) {
	if sal < 15000 {
		return "", &Error{"O salário digitado não está dentro do valor mínimo"}
	}
	return "Necessário pagamento de imposto", nil
}
