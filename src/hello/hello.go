package main

import "fmt"
import "reflect"

func main() {
	//Declaração completa com a definição explícita dos tipos
	//var nome string = "Ivo Rafael"
	//var idade int = 39
	//var versao float32 = 1.1

	// Ou

	//Declaração com inferência de tipos pela própria linguagem
	//var nome = "Ivo Rafael"
	//var idade = 39
	//var versao = 1.1

	//Ou
	//Atribuição de forma curta e com inferência de tipos pela própria linguagem
	nome := "Ivo Rafael"
	idade := 39
	versao := 1.1

	fmt.Println("Olá Sr.", nome, "sua idade é de", idade, "anos.")
	fmt.Println("Este programa está na versão: ", versao)

	fmt.Println("O tipo da variável nome é:", reflect.TypeOf(nome))
	fmt.Println("O tipo da variável idade é:", reflect.TypeOf(idade))
	fmt.Println("O tipo da variável versao é:", reflect.TypeOf(versao))
}
