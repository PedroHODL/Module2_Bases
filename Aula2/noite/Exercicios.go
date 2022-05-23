package main

import "fmt"

func main() {
	Ex_1()
	fmt.Println("--------------Ex 2-------------------")
	Ex_2()
}

func Ex_1() {
	aluno1 := Estudante{}
	aluno1.addEstudante("Pedro", "Henrique", "12.345.678", "16/05/2022")
	aluno1.printEstudante()
}

type Estudante struct {
	Nome         string
	Sobrenome    string
	RG           string
	DataAdmissao string
}

func (e *Estudante) addEstudante(nome, sobrenome, rg, data string) {
	e.Nome = nome
	e.Sobrenome = sobrenome
	e.RG = rg
	e.DataAdmissao = data
}

func (e Estudante) printEstudante() {
	fmt.Println("O aluno", e.Nome+" "+e.Sobrenome, "com RG", e.RG,
		"foi matriculado no dia", e.DataAdmissao)
}

// ------------------------------ Ex 2 ------------------------------
func Ex_2() {
	produto1 := novoProduto("Grande", "Geladeira", 10000)
	produto2 := novoProduto("Pequeno", "Smartphone", 4000)
	produto3 := novoProduto("Médio", "Cadeira Gamer", 15000)
	
	loja := novaLoja()
	loja.Adicionar(produto1)
	loja.Adicionar(produto2)
	loja.Adicionar(produto3)

	fmt.Printf("O total da compra é de $%.2f", loja.Total())
}

func novoProduto(tipo, nome string, preço float32) produto {
	return produto{Tipo: tipo, nome: nome, preço: preço}
}

type produto struct {
	Tipo  string
	nome  string
	preço float32
}

type Produto interface {
	CalcularCusto() float32
}

func (p produto) CalcularCusto() float32 {
	var preco float32

	switch p.Tipo {
	case "Pequeno":
		preco = p.preço
	case "Médio":
		preco = p.preço * 1.03
	case "Grande":
		preco = p.preço*1.06 + 2500
	}
	return preco
}

type loja struct {
	listaProdutos []produto
}

type ECommerce interface {
	Total() float32
	Adicionar() error
}

func (l loja) Total() float32 {
	var total float32 = 0
	for _, produto := range l.listaProdutos {
		total += produto.CalcularCusto()
	}
	return total
}

func (l *loja) Adicionar(p produto) {
	l.listaProdutos = append(l.listaProdutos, p)
}

func novaLoja() loja {
	var l loja
	return l
}