package loops

import (
	"curso_basico_go/modules/functions"
	"fmt"
)

func ForLoops() {
	// Conditional for
	for i := 0; i <= 10; i++ {
		functions.MyPrint(i)
	}

	fmt.Println()

	// For While
	counter := 0
	for counter < 10 {
		functions.MyPrint(counter)
		counter++
	}

	fmt.Println()

	// For forever
	counterForever := 0
	for {
		functions.MyPrint(counterForever)
		counterForever++
	}
}
