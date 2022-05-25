package Ex2

import "fmt"

func Ex2() {
	user1 := novoUser("Pedro", "Henrique", "gomeligo@meli.com")
	geladeira := novoProduto("Geladeira", 4000.0)
	mangas := novoProduto("Mangás", 34.90)

	user1.addProduto(&geladeira, 1)
	user1.addProduto(&mangas, 5)

	user1.printarProdutos()
}

type Usuario struct {
	Nome      string
	Sobrenome string
	Email     string
	Produtos  []Produto
}

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

func novoUser(nome, sobrenome, email string) Usuario {
	return Usuario{nome, sobrenome, email, []Produto{}}
}

func novoProduto(nome string, preco float64) Produto {
	return Produto{nome, preco, 0}
}

func (u *Usuario) addProduto(prod *Produto, qtd int) {
	prod.Quantidade = qtd
	u.Produtos = append(u.Produtos, *prod)
}

func (u *Usuario) deleteProduto() {
	u.Produtos = []Produto{}
}

func (u Usuario) printarProdutos() {
	fmt.Printf("Lista de compras de %s\n", u.Nome+" "+u.Sobrenome)
	for _, prod := range u.Produtos {
		fmt.Printf("Produto: '%s', custa R$%.2f e irei comprar %dx\n",
			prod.Nome, prod.Preco, prod.Quantidade)
	}
}
