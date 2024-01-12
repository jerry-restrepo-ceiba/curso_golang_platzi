package main

import "fmt"

func main() {
	const pi float64 = 3.14
	const pi2 = 3.15

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	suma := x + y
	fmt.Println("suma", suma)

	resta := x - y
	fmt.Println("resta", resta)

	multipliacion := x * y
	fmt.Println("multipliacion", multipliacion)

	division := x / y
	fmt.Println("division", division)

	modulo := x % y
	fmt.Println("modulo", modulo)

	x++
	fmt.Println("incremental", x)

	x--
	fmt.Println("decrementa", x)

	//Paquete fmt

	helloMessage := "hello"
	worldMessage := "world"

	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//printf
	nombre := "platzi"
	cursos := 500

	fmt.Printf("%s tiene mas de %d cursos \n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos \n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos \n", nombre, cursos)
	fmt.Println("message: ", message)

	// tipos de datos
	fmt.Printf("helloMessage: %T \n\n", helloMessage)
	fmt.Printf("CURSOS: %T \n", cursos)
}

/*

 */
