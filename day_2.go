// Golang é uma linguagem fortemente tipada, cada variável tem seu tipo e não pode mudar.
// Inteiros (int, int8, int16, int32, int64)
// Floats (float32, float64)
// Booleanos (bool)
// Strings (string)

// Declaração de variáveis
// Longa(var)
// var idade int = 25
// var nome string = "Carlos"
// var altura float64 = 1.75
// var ativo bool = true

// Declaração Curta (Usando :=)
// É usada apenas dentro de funções e o tipo é inferido automaticamente.
// idade := 25
// nome := "Carlos"
// altura := 1.75
// ativo := true

// Constantes são valores que não podem ser alterados após a atribuição
// const PI = 3.14159
// const Nome = "Golang"

package main

import "fmt"

var idade int = 25
var altura float64 = 1.75
var nome string = "Igor"

const PI = 3.14

func main() {
	humano := 1
	fmt.Println("minha idade é:", idade)
	fmt.Println("minha altura é:", altura)
	fmt.Println("meu nome é:", nome)
	fmt.Println("humano:", humano)
	fmt.Println("PI:", PI)
}
