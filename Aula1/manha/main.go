package main

import "fmt"

func main() {
	//Ex-1 - Imprimindo o nome na tela
	/**
		1. Crie uma aplicação que tenha uma variável “nome” e outra “idade”.
		2. Imprima no terminal o valor de cada variável.
	**/
	var nome string = "Pedro"
	var idade int = 20
	fmt.Println("\nMeu nome é", nome, "e tenho", idade, "anos")

	//Ex-2 - Clima
	/**
		Uma empresa de meteorologia quer ter um sistema onde possa ter a temperatura, umidade e
		pressão atmosférica de diferentes lugares.

		1. Declare 3 variáveis especificando o tipo de dado delas, como valor elas devem
			possuir a temperatura, umidade e pressão atmosférica de onde você se encontra.
		2. Imprima os valores no console.
		3. Quais tipos de dados serão atribuídos a essas variáveis?
	**/
	var temperatura int = 9
	var umidade int = 71
	var pressao int = 1017
	fmt.Print("\nA temperatura é de ", temperatura, " graus",
		"\na umidade do ar é de ", umidade, "%",
		"\ne a pressão é de ", pressao, "mbar\n\n")

	//Ex-3 - Declaração de variáveis
	/**
	Um professor de programação está corrigindo as avaliações de seus alunos na disciplina de
	Programação I para fornecer-lhes o feedback correspondente. Um dos pontos do exame é
	declarar diferentes variáveis.

	Ajude o professor com as seguintes questões:
		1. Verifique quais dessas variáveis declaradas pelo aluno estão corretas.
		2. Corrigir as incorrectas.

			var 1nome string
			var sobrenome string
			var int idade
			1sobrenome := 6
			var licenca_para_dirigir = true
			var estatura da pessoa int
			quantidadeDeFilhos := 2

		Corrigido:
			var nome1 string
			var sobrenome string
			var idade int
			sobrenome1 := 6
			var licenca_para_dirigir bool = true
			var estatura_da_pessoa int
			quantidadeDeFilhos := 2
	**/

	//Ex-4 - Tipos de dados
	/**
	Um estudante de programação tentou fazer declarações de variáveis com seus respectivos
	tipos em Go mas teve várias dúvidas ao fazê-lo. A partir disso, ele nos deu seu código e
	pediu a ajuda de um desenvolvedor experiente que pode:

		● Revisar o código e realizar as devidas correções.

			var sobrenome string = "Silva"
			var idade int = "25"
			boolean := "false";
			var salario string = 4585.90
			var nome string = "Fellipe"

		Corrigido:
			var sobrenome string = "Silva"
			var idade int = 25
			boolean := false
			var salario float32 = 4585.90
			var nome string = "Fellipe"
	**/

}
