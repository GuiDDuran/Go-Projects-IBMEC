package main

import (
	"bufio"
	"fmt"
	"os"
)

/* 01. Elabore um struct Contato, que deve conter os campos nome e e-mail. O struct deve conter um método para alterar o e-mail. */

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) alteraEmail(novoEmail string) {
	c.Email = novoEmail
}

/* 02. Elabore uma função adicionarContato, que recebe um contato e um array com até 5 elementos e inclui o contato no primeiro elemento vazio do array. */

func adicionarContato(contato Contato, contatos *[5]Contato) {
	for i := range contatos {
		if contatos[i].Nome == "" && contatos[i].Email == "" {
			contatos[i] = contato
			return
		}
	}
}

/* 03. Elabore uma função excluirContato, que recebe um array de 5 elementos e elimina o último contato válido. */

func excluirContato(contatos *[5]Contato){
	for i := len(*contatos) - 1; i >= 0; i-- {
		if contatos[i].Nome != "" && contatos[i].Email != "" {
			contatos[i] = Contato{}
			return
		}
	}
}

func main() {
	var contatos [5]Contato
	leitor := bufio.NewReader(os.Stdin)

	/* 04. Elabore uma interface por linha de comando na função main, que cria um array de 5 elementos e permite a inclusão ou exclusão de contatos. */
	for{
		fmt.Println("1. Adicionar contato.")
		fmt.Println("2. Excluir último contato.")
		fmt.Println("3. Encerrar o programa.")
		fmt.Println("Opção: ")
		opcao, _ := leitor.ReadString('\n')

		switch string(opcao[0]){
		case "1":
			var nome, email string

			fmt.Println("Digite o nome do contato: ")
			nome, _ = leitor.ReadString('\n')
			fmt.Println("Digite o email do contato: ")
			email, _ = leitor.ReadString('\n')

			novoContato := Contato{Nome: nome, Email: email}
			adicionarContato(novoContato, &contatos)

		case "2":
			excluirContato(&contatos)

		default:
			fmt.Println("Opção inválida.")
		}
	}
}
