package Ex1

import (
	"fmt"
	"strings"
)

func Ex1() {
	us := criarUsuario("Pedro", "Henrique Oliveira", 19, "peter@meli.com", "1234")

	// Antes de alterar
	fmt.Println(us.Nome+" "+us.Sobrenome, "tem", us.idade, "anos\nemail:",
		us.email, "\nsenha:", us.senha)

	us.setName("Pedro Henrique Oliveira Dantas Lopes")
	us.setIdade(20)
	us.setEmail("pedro@meli.com")
	us.setSenha("SenhaSegura123")

	fmt.Println("\nApós alterações:")
	fmt.Println(us.Nome+" "+us.Sobrenome, "tem", us.idade, "anos\nemail:",
		us.email, "\nsenha:", us.senha)
}

type Usuario struct {
	Nome      string
	Sobrenome string
	idade     int
	email     string
	senha     string
}

func (u *Usuario) setName(n string) {
	nome := strings.Split(n, " ")
	u.Nome = nome[0]
	u.Sobrenome = strings.Join(nome[1:], " ")
}

func (u *Usuario) setIdade(num int) {
	u.idade = num
}

func (u *Usuario) setEmail(e string) {
	u.email = e
}

func (u *Usuario) setSenha(s string) {
	u.senha = s
}

func criarUsuario(nome, sobrenome string, idade int, email, senha string) Usuario {
	return Usuario{nome, sobrenome, idade, email, senha}
}
