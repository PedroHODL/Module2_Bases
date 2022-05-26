package Exercicios

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const minID int = 100000
const maxID int = 999999

type Cliente struct {
	Arquivo  int
	Nome     string
	RG       string
	Telefone int
	Endereco string
}

func erro() {
	err := recover()

	if err != nil {
		fmt.Println("Error:", err)
	}
}

func Ex2() {
	defer erro()

	user := novoCliente("PedroH", "123.456.789-0", 912345678, "Rua aleatória")
	fmt.Println(user)
}

func novoCliente(nome, rg string, telefone int, endereco string) Cliente {
	if nome == "" || rg == "" || telefone == 0 || endereco == "" {
		panic("algum dado preenchido é nulo")
	}

	novoC := Cliente{
		0,
		nome,
		rg,
		telefone,
		endereco,
	}

	clientes := abrirCsv("./GoBases4/Aula2/costumers.csv")
	for _, cliente := range clientes {
		temp := cliente.Arquivo
		cliente.Arquivo = 0

		if cliente == novoC {
			panic("o cliente já existe")
		}
		cliente.Arquivo = temp
	}

	novoC.Arquivo = gerarIdArquivo()

	return novoC
}

func gerarIdArquivo() int { // Função que gera um ID de 6 digitos
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxID-minID) + minID
}

func abrirCsv(local string) []Cliente {
	arquivo, err := os.Open(local)
	if err != nil {
		panic(fmt.Sprint("caminho incorreto ou arquivo inexistente"))
	}
	defer arquivo.Close()

	csvLinhas, err := csv.NewReader(arquivo).ReadAll()
	if err != nil {
		panic(fmt.Sprint("separador incorreto ou arquivo csv inválido"))
	}

	var listaClientes []Cliente
	for _, cliente := range csvLinhas {
		id, _ := strconv.Atoi(cliente[0])
		tel, _ := strconv.Atoi(cliente[3])
		user := Cliente{id, cliente[1], cliente[2], tel, cliente[4]}

		listaClientes = append(listaClientes, user)
	}

	return listaClientes
}
