package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415 //Forma padrão de definição de constantes
	var raio = 3.2

	//Utilizar o := é uma forma de declarar e inicializar uma variavel
	area := PI * math.Pow(raio, 2) //Basicamente colocando o valor do raio ao quadrado
	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	//Podemos declarar variáveis de maneira simultanea
	var e, f bool = true, false

	fmt.Println(e, f)

	//Podemos declarar variáveis de maneira simultanea

	g, h, i := 2, false, "epa!"

	fmt.Println(g, h, i)

}
