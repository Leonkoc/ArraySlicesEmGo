package main

import "fmt"

func main() {
	nomeDoUsuario := ""
	sobrenomeUsuario := ""
	emailUsuario := ""
	numeroDeTickets := ""

	fmt.Println("Digite o seu nome: ")
	fmt.Scan(&nomeDoUsuario)

	fmt.Println("E seu sobrenome: ")
	fmt.Scan(&sobrenomeUsuario)

	println("Digite um email de contato: ")
	fmt.Scan(&emailUsuario)

	println("Quantos tikets para conferencia você deseja?")
	fmt.Scan(&numeroDeTickets)

	var pessoasComReserva = []string{}
	pessoasComReserva = append(pessoasComReserva, "André")
	pessoasComReserva = append(pessoasComReserva, "João")
	pessoasComReserva = append(pessoasComReserva, nomeDoUsuario+" "+sobrenomeUsuario)

	fmt.Printf("A lista de pessoas com reservas garantidas %v\n", pessoasComReserva)

}
