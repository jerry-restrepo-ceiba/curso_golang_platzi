package mypackage

import "fmt"

// CarPublic con acce publico
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

//PrintMessage imprimir privado
func PrintMessage(message string) {
	fmt.Println("Hola aqui ", message)
	printMessagePrivate("juan")
}

func printMessagePrivate(message string) {
	fmt.Println("Hola privado ", message)
}
