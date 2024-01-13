package main

import (
	pk "curso_golang_platzi/src/mypackage"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo"
)

// colección de campos, estructura de datos
type car struct {
	brand string
	year  int
}

type pc struct {
	ram   int
	disk  int
	brand string
}

type cuadrado struct {
	base float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

type rectangulo struct {
	base   float64
	altura float64
}

type figura2D interface {
	area() float64
}

func (r rectangulo) areaRectangulo() float64 {
	return r.base * r.altura
}

func calcular(f figura2D) {
	fmt.Println("Area: ", f.area())
}

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func saychan(text string, c chan<- string) {
	c <- text
}

func message(text string, c chan string) {
	c <- text
}

func main() {

	//clasesPrimeras()

	//clasesFunciones()

	//clasesCiclosForCondicional()

	//clasesIfCondicional(1, 2)

	//clasesSwitch(8)

	//claseUsoKeywordDeferBreakConinue()

	//clasesArraySlices()

	//clasesArraysSlicesArrange()

	//claseLlaveValorMaps()

	//claseStructs()

	//claseModificadorAcceso()

	//claseStructspunteros()

	//clasesStringers()

	//claseInterfaceListaInterface()

	//claseunoGoroutines()

	//clasesChannels()
	//claseRange()

	claseModulos()

}

func claseModulos() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "HELLO WORLD")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func claseRange() {
	c := make(chan string, 2)
	c <- "mensaje 1"
	c <- "mensaje 2"

	fmt.Println(len(c), cap(c))

	//range y close

	close(c) // cerrar canal no va a recibir ningun otro dato adicional

	for message := range c {
		fmt.Println(message)
	}

	//select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje 1.1", email1)
	go message("mensaje 2.1", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("email recibido de email1: ", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2: ", m2)
		}
	}
}

func clasesChannels() {
	c := make(chan string, 1)

	fmt.Println("hello")

	go saychan("bye", c)

	fmt.Println(<-c)
}

func claseunoGoroutines() {
	var wg sync.WaitGroup
	fmt.Println("hello")
	wg.Add(1)
	go say("world", &wg)
	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second * 1)
}

func claseInterfaceListaInterface() {
	mycuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 3, altura: 4}

	calcular(mycuadrado)
	calcular(myRectangulo)

	//listaIinterfaces

	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)

}

func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

func clasesStringers() {
	myPc := pc{ram: 16, brand: "msi", disk: 100}

	fmt.Println(myPc)
}

func (myPC pc) ping() {
	fmt.Println("my", myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func claseStructspunteros() {
	a := 50
	b := &a
	fmt.Println(b)

	//* acceder al valor de memoria
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	a = 200
	fmt.Println(*b)

	myPC := pc{ram: 16, disk: 200, brand: "msi"}

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)

}

func claseModificadorAcceso() {
	// en funciones y structs primera letra sea en mayuscula publico sino privado
	var myCar pk.CarPublic
	myCar.Brand = "ferrari"
	myCar.Year = 2025
	fmt.Println(myCar)
	pk.PrintMessage("jerry")

}

func claseStructs() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// otra manera
	var otherCar car
	otherCar.brand = "Ferrari"

	fmt.Println(otherCar)
}

func claseLlaveValorMaps() {
	m := make(map[string]int)
	m["jose"] = 14
	m["pepito"] = 20

	fmt.Println(m)

	for i, v := range m {
		fmt.Println(i, v)
	}

	// encontrar un valor
	value, ok := m["jose"]
	fmt.Println(value, ok)

}

func clasesArraysSlicesArrange() {
	slice := []string{"hola", "que", "hace"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	for _, valor := range slice {
		fmt.Println("solo necesito valor", valor)
	}

	for i := range slice {
		fmt.Println("solo necesito el indice ", i)
	}

	isPalindromo("amor a roma")
}

func isPalindromo(value string) {
	var textReverse string

	for i := len(value) - 1; i >= 0; i-- {
		textReverse += string(value[i])
	}

	fmt.Println(textReverse)
	fmt.Println((value == textReverse))
}

func clasesArraySlices() {
	var array [4]int // valor inmutable no se puede agregar otro elelemnto pero si modificar

	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	//slice  diferencia no se le indica la cantidad de vbalore a tener

	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//apened

	slice = append(slice, 7)
	fmt.Println(slice, len(slice), cap(slice))

	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice, len(slice), cap(slice))
}

func claseUsoKeywordDeferBreakConinue() {
	//defer ejecuta la ultima funcion antes de morir  uso de abrir archivos arir base de datos
	//cerrar conexiones
	//defer por funcion
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//continue y break se utilian en for

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		//continue error controlado y ejecutar el ciclo for
		if i == 2 {
			fmt.Println("es dos")
			continue
		}

		// break
		if i == 6 {
			fmt.Println("encontro 6")
			break
		}

	}
}

func clasesSwitch(a int) {
	modulo := a % 2
	switch modulo {
	case 0:
		fmt.Println("es par")
	default:
		fmt.Println("es impar")
	}

	//sin condicion cuando quieres anidar multiples condiciones

	value := 50
	switch {
	case value > 100:
		fmt.Println("es mayor a 100")
	case value < 0:
		fmt.Println("es menor a 0")
	default:
		fmt.Println("no condicion")
	}

}

func clasesIfCondicional(valor1, valor2 int) {

	if valor1 == 1 {
		fmt.Println("es uno")
	} else {
		fmt.Println("no es uno")
	}

	// with and

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("and es verdad")
	}

	//with or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("or es verdad")
	} else if valor1 == 3 {
		fmt.Println("or es verdad 3")
	}

	// convertir texto a numero
	value, err := strconv.Atoi("adas53")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("es tipo %T", value)

}

func clasesCiclosForCondicional() {
	for i := 0; i <= 10; i++ {
		fmt.Println("entro al ciclo for: ", i, "veces")
	}

	//for while
	counter := 0
	for counter < 10 {
		fmt.Println("entro al while", counter)
		counter++
	}

	// for forever
	counterForever := 0
	for {
		fmt.Println("infinito", counterForever)
		counterForever++
	}

	//  for
}

func clasesFunciones() {
	normalFunction()
	normalFunctionParameter("message string")
	value := returnValue(2)
	fmt.Println(value)
	value1, value2 := doubleReturn(2)
	fmt.Println("value1 y value 2", value1, value2)
	value3, _ := doubleReturn(2)
	fmt.Println("value3", value3)
}

func clasesPrimeras() {
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

func normalFunction() {
	fmt.Println("Hola Mundo")
}

func normalFunctionParameter(message string) {
	fmt.Println(message)
}

func multipleParameterFunction(a, b, c int, message string) {
	fmt.Println("multiple parametros", a, b, c, message)
}

func returnValue(a int) int {
	return a + 2
}

func doubleReturn(a int) (c, d int) {
	return a + 1, a * 2
}
