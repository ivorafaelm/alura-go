package main

import "fmt"
import "os"
import "net/http"
import "time"
import "bufio"
import "strings"
import "io"

const monitoramento = 5
const delay = 10

func main() {

	exibeIntroducao()
	for {
		exibeMenu()

		comando := leComando()
		//fmt.Scanf("%d", &comando)
		//Ou
		//fmt.Scan(&comando) //Neste caso a função já sabe que o valor que pode ser recebido é do tipo inteiro, pois é o que foi declarado
		//fmt.Println("O endereço da minha variável comando é:", &comando)
		//fmt.Println("O comando escolhido foi", comando)

		//if comando == 1 {
		//	fmt.Println("Monitorando ...")
		//} else if comando == 2 {
		//	fmt.Println("Exibindo Logs ...")
		//} else if comando == 0 {
		//	fmt.Println("Saindo do programa")
		//} else {
		//	fmt.Println("Não conheço este comando")
		//}

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs ...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0) //Informa ao SO que o programa está sendo encerrado da forma esperada.
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1) //Informa ao SO que o programa não sendo encerrado da forma inesperada - com erro.
		}

	}

}

func exibeIntroducao() {
	nome := "Ivo Rafael"
	versao := 1.1

	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa está na versão: ", versao)
}

func exibeMenu() {
	fmt.Println("1-Iniciar Monitoramento")
	fmt.Println("2-Exibir Logs")
	fmt.Println("0-Sair do Programa")
}

func leComando() int {
	var comandoLido int
	//fmt.Scanf("%d", &comando)
	//Ou
	fmt.Scan(&comandoLido) //Neste caso a função já sabe que o valor que pode ser recebido é do tipo inteiro, pois é o que foi declarado
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando ...")

	//sites := []string{"https://www.google.com", "https://www.globo.com", "https://www.uol.com.br"}

	sites := leSitesDoArquivo()

	fmt.Println(sites)

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando o site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code retornado foi:", resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')

		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()
	return sites
}
