package Ex3

import "fmt"

func Ex3() {
	p := []Produto{{
		Nome:       "Geladeira",
		Preco:      2000.0,
		Quantidade: 2,
	}, {
		Nome:       "Mangás",
		Preco:      100.00,
		Quantidade: 5,
	}}

	s := []Serviço{{
		Nome:  "Instalar antena",
		Preco: 8000.0,
		Mins:  45,
	}}

	m := []Manutenção{{
		Nome:  "Arrumar janela",
		Preco: 1000.0,
	}}

	c := make(chan float64)
	go somarProdutos(p, c)
	go somarServicos(s, c)
	go somarManutencao(m, c)

	total := 0.0
	for i := 0; i < 3; i++ {
		total += <-c
	}

	fmt.Printf("O total do serviço completo é de R$%.2f\n", total)

}

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

type Serviço struct {
	Nome  string
	Preco float64
	Mins  int
}

type Manutenção struct {
	Nome  string
	Preco float64
}

func somarProdutos(array []Produto, c chan float64) {
	total := 0.0
	for _, prod := range array {
		total += prod.Preco * float64(prod.Quantidade)
	}
	c <- total
}

func somarServicos(array []Serviço, c chan<- float64) {
	total := 0.0
	for _, serv := range array {
		total += serv.Preco * (float64(serv.Mins) / 60.0)
	}
	c <- total
}

func somarManutencao(array []Manutenção, c chan<- float64) {
	total := 0.0
	for _, manu := range array {
		total += manu.Preco
	}
	c <- total
}
