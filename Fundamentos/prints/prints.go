package main

import "fmt"

func main() {
	//Coloca tudo na mesma linha
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	//Quebra linha
	fmt.Println(" Nova")
	fmt.Println(" linha.")

	x := 3.141516

	//Aqui deu erro pois queriamos concatenar um number com uma string
	//fmt.Println("O valor de x é" + x)

	//Basicamente criamos uma váriavel com o metodo Sprint e passamos o x como parametro
	//Isso faz com que retorne uma string a partir do valor do parametro
	xs := fmt.Sprint(x)

	//Com isso, retornamos o valor de x
	fmt.Println("O valor de x é " + xs)

	//Outra maneira de realizar o mesmo sem criar uma nova variável
	fmt.Println("O valor de x é", x)

	//Podemos formatar nosso print
	fmt.Printf("O valor de x é %.2f", x)

	//Mais 1 exemplo de uso

	a := 1
	b := 1.9999
	c := false
	d := "oi"

	// \n para quebrar linha
	// %d = inteiro, %f = float, %.1f = float com 1 casa, %t = boolean, %s = string
	// existe o %v = que imprime muitos tipos diferentes, como se fosse genêrico
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

}
