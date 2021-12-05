package main

import "fmt"

func main() {

	nome := "Ivo Rafael"
	versao := 1.1

	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa está na versão: ", versao)

	fmt.Println("1-Iniciar Monitoramento")
	fmt.Println("2-Exibir Logs")
	fmt.Println("0-Sair do Programa")

	var comando int
	//fmt.Scanf("%d", &comando)
	//Ou
	fmt.Scan(&comando) //Neste caso a função já sabe que o valor que pode ser recebido é do tipo inteiro, pois é o que foi declarado
	fmt.Println("O endereli da minha variável comando é:" & comando)
	fmt.Println("O comando escolhido foi", comando)

}
