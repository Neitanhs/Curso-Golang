package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	//números ineiros
	fmt.Println(1, 2, 1000)
	//reflect é uma propriedade que retorna em terminal o tipo do dado inserido
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//Inteiros apenas positivos --> uint8, uint16, uint32 e uint64

	var b byte = 255 //Esse byte é um apelido para o uint8 (nem todos possuem)
	fmt.Println("O byte é", reflect.TypeOf(b))

	//Inteiros com sinal --> int8, int16, int32 e int64

	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	//Alias Rune representa um mapeamento da tabela Unicode (int32)
	var i2 rune = 'a'
	fmt.Println("rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	//Números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))
	//Isso quer dizer que por padrão números fracionados são float64

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	//E por ser uma variável boleana podemos então realizar alterações lógicas
	fmt.Println(!bo) //False

	//String
	s1 := "Olá"
	fmt.Println(s1 + "!") //Concatenação
	fmt.Println("O tamanho da string é", len(s1))

	//String de multiplas linhas
	s2 := `Oá
	meu
	nome
	é
	Nathan`
	fmt.Println("O tamanho da String é", len(s2))

	//char
	char := 'a'                                           //Aspa única
	fmt.Println("O tipo de char é", reflect.TypeOf(char)) //Não existe tipo char

}
