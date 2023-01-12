package main

// Quando você quer que seja um executável , esse package precisa ser declarado Main

import "fmt"
import "reflect"

// import de fmt - " formatação" - Para usar o println

//criando uma função main como no proprio java, mas a declaração é bem mais simples
func main() {
	var nome string = "Marco"
	var idade int = 34
	var versao float32 = 1.1

	nome := "Marco"
	idade := 34
	versao := 1.1

	// nome := "marco"
	// Não é necessário colocar o var indicando que foi criada a variável, só precisa colocar o :=
	//e o Tipo também é dinamico, tomar cuidado em exemplos de float 32/64, sempre que necessário ele irá colocar a maior de todas





	// variaveis declaradas e não utilizadas vão gerar erro e o código não será nem compilado
	fmt.Println("Sua idade é:",idade)
	fmt.Println("Olá! Tudo bem?", nome)
	fmt.Println("A versão é :", versao)
	fmt.Println("O tipo de nome é:", reflect.TypeOf(nome))
	fmt.Println("O tipo de idade é:", reflect.TypeOf(idade))
	fmt.Println("O tipo de versao é:", reflect.TypeOf(versao))
	


	fmt.Println("Olá mundo")
}

// para executar o programa, ir pelo terminal na pasta e escrever
// go build nomem doarquivo.go
// para compilar e depois executar

// para rodar diretamente pelo terminal, onde ele compila e ja roda
// -- go run nomedoarquivo.go
