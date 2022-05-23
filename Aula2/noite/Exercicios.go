package main

import "fmt"

func main() {
	Ex1()
	fmt.Println("--------------Ex 2-------------------")
	Ex2()
}

func Ex1() {
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

// Ex2 ------------------------------ Ex 2 ------------------------------
func Ex2() {
	produto1 := novoProduto("Grande", "Geladeira", 10000)
	produto2 := novoProduto("Pequeno", "Smartphone", 4000)
	produto3 := novoProduto("Médio", "Cadeira Gamer", 15000)

	fmt.Printf("O preço do produto1 é $%.2f \n", produto1.CalcularCusto())
	loja := novaLoja()
	loja.Adicionar(produto1)
	loja.Adicionar(produto2)
	loja.Adicionar(produto3)

	fmt.Printf("O total da compra é de $%.2f \n", loja.Total())
}

func novoProduto(tipo, nome string, preco float32) Produto {
	var newProduct = produto{tipo, nome, preco}
	return &newProduct
}

type produto struct {
	Tipo  string
	nome  string
	preco float32
}

type Produto interface {
	CalcularCusto() float32
}

func (p produto) CalcularCusto() float32 {
	var preco float32

	switch p.Tipo {
	case "Pequeno":
		preco = p.preco
	case "Médio":
		preco = p.preco * 1.03
	case "Grande":
		preco = p.preco*1.06 + 2500
	}
	return preco
}

type loja struct {
	listaProdutos []Produto
}

type ECommerce interface {
	Total() float32
	Adicionar(p Produto)
}

func (l loja) Total() float32 {
	var total float32 = 0
	for _, produto := range l.listaProdutos {
		total += produto.CalcularCusto()
	}
	return total
}

func (l *loja) Adicionar(p Produto) {
	l.listaProdutos = append(l.listaProdutos, p)
}

func novaLoja() ECommerce {
	return &loja{}
}
