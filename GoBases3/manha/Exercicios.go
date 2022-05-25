package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	loja := novaLoja()
	loja.Adicionar(novoProduto(111223, 30012.00, 1))
	loja.Adicionar(novoProduto(444321, 1000000.00, 4))
	loja.Adicionar(novoProduto(434321, 50.50, 1))

	loja.GerarCsv()
	ExibirCsv("./GoBases3/manha/produtos.csv")
	loja.PullInfoCsv("./GoBases3/manha/produtos.csv")
}

func (l loja) GerarCsv() {
	infoCsv := []byte("ID;Preço;Quantidade\n")
	for _, prod := range l.listaProdutos {
		res := fmt.Sprintf("%d;%.2f;%d\n", prod.ID, prod.Preco, prod.Quantidade)
		infoCsv = append(infoCsv, []byte(res)...)
	}

	err := os.WriteFile("./GoBases3/manha/produtos.csv", infoCsv, 0644)
	if err != nil {
		fmt.Print("Erro ao escrever no arquivo:")
		panic(err)
	}
}

func (l *loja) PullInfoCsv(local string) {
	csvLinhas := abrirCsv(local)

	for _, linha := range csvLinhas {
		id, _ := strconv.ParseFloat(linha[0], 64)
		preco, _ := strconv.ParseFloat(linha[1], 64)
		qtd, _ := strconv.ParseFloat(linha[2], 64)
		l.Adicionar(novoProduto(int(id), preco, int(qtd)))
	}
}

func ExibirCsv(local string) {
	csvLinhas := abrirCsv(local)

	total := 0.0
	for _, linha := range csvLinhas {
		fmt.Printf("%-6s%15s%15s\n", linha[0], linha[1], linha[2])

		preco, _ := strconv.ParseFloat(linha[1], 64)
		qtd, _ := strconv.ParseFloat(linha[2], 64)
		total += preco * qtd
	}
	fmt.Printf("%21.2f", total)
}

func abrirCsv(local string) [][]string {
	arquivo, err := os.Open(local)
	if err != nil {
		fmt.Print("Erro ao abrir o arquivo:")
		panic(err)
	}
	defer arquivo.Close()

	temp := csv.NewReader(arquivo)
	temp.Comma = ';'
	csvLinhas, err := temp.ReadAll()
	if err != nil {
		fmt.Print("Erro na divisão dos indices:")
		panic(err)
	}
	return csvLinhas
}

type produto struct {
	ID         int
	Preco      float64
	Quantidade int
}

func novoProduto(id int, preco float64, quantidade int) produto {
	var newProduct = produto{id, preco, quantidade}
	return newProduct
}

type loja struct {
	listaProdutos []produto
}

func (l loja) Total() float64 {
	total := 0.0
	for _, produto := range l.listaProdutos {
		total += produto.Preco * float64(produto.Quantidade)
	}
	return total
}

func (l *loja) Adicionar(p produto) {
	l.listaProdutos = append(l.listaProdutos, p)
}

func novaLoja() loja {
	return loja{}
}
