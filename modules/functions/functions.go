package functions

import "fmt"

func MyPrint(message ...interface{}) {
	fmt.Println(message...)
}

func ftmPackage() {
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene más de %d cursos\n", nombre, cursos)

	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Type
	fmt.Printf("Name: %T\n", nombre)
}

func tripeArguments(a, b int, c string) {
	MyPrint(a, b, c)
}

func Pow(a int) int {
	return a * a
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func AllFunctions() {
	ftmPackage()
	MyPrint("Hola")
	tripeArguments(1, 2, "3")
	MyPrint(Pow(5))

	v1, v2 := doubleReturn(2)
	v3, _ := doubleReturn(32)
	MyPrint(v1, v2, v3)

}
