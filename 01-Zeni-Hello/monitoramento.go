package main

import "fmt"

func main() {
	nome := "Marco Castellani"
	versao := 1.0

	fmt.Println("Olá ",nome)
	fmt.Println("O programa está na versão ",versao)

	fmt.Println("1 - Monitorar os Sites")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")

	var comando int
	fmt.Scanf("%d",&comando)

	/* -------------Sobre OS SCAN
	//& é o endereço da variável comando
	//%d ´o modificador da variável, o que vai ser recebido
	// um println de comando gera o que tem dentro da variável							---- valor
	// ja o println de &comando mostra o local, "0x00909e090" , exemplo de posição		---- endereço

	//fmt.scan(&comando)
	// nesse caso o Scan entende como a var é do tipo int, ele precisa receber um Int, 
	// você apenas precisa declarar a variável, onde vai ser inserido o valor
	// Caso seja jogada uma string, ou coisa que não é esperada para a variável em especifica
	// ele vai retornar o valor 0, como tda variável é Zero


	fmt.Println(comando)
	*/

	/*  ----------- IF
	if comando == 1{ // o comando If só aceita um teste , que retorne True or False
		fmt.Println("Monitorando os Sites")
	} else if comando ==2 {
		
		fmt.Println("Exibindo Logs")
	} else if comando == 0 {
		
		fmt.Println("Saindo do Programa")
	}else {
		fmt.Println("Não inseriu comando esperado.")
	}
	*/
	
	switch comando {
		// switch não precisa de um break no final de cada Case. 
		// Ele apenas entra no case e pula fora quando termina 
	case 1:
		fmt.Println("Monitorando os Sites")
	case 2:
		fmt.Println("Exibindo Logs")
	case 0:
		fmt.Println("Saindo do Programa")
	default:
		fmt.Println("Não inseriu comando esperado.")
	}



}