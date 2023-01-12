package main

import "fmt"      // pacote que faz a formatação
import "os"       // pacote que fala diretamente com o Sistema operacional
import "net/http" // Pacote que acessa a internet e faz requisições

func main() {
	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	switch comando {
	// switch não precisa de um break no final de cada Case.
	// Ele apenas entra no case e pula fora quando termina
	case 1:
		iniciaMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs")
	case 0:
		fmt.Println("Saindo do Programa")
		os.Exit(0)
	default:
		fmt.Println("Não inseriu comando esperado.")
		os.Exit(-1) // para dizer que deu alguma coisa de errada, e não só fechou o programa
	}
}

func exibeIntroducao() {
	nome := "Marco Castellani"
	versao := 1.4

	fmt.Println("Olá ", nome)
	fmt.Println("O programa está na versão ", versao)
}
func exibeMenu() {
	fmt.Println("1 - Monitorar os Sites")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	var comandolido int
	fmt.Scan(&comandolido)
	return comandolido
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando os Sites...")
	site := "https://www.google.com.br"
	resp := http.Get(site)
}
