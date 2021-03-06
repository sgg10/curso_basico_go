package pointers

import (
	"curso_basico_go/modules/functions"
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	functions.MyPrint(myPc.brand, "Pong")
}

func (myPc *pc) duplicateRAM() {
	myPc.ram *= 2
}

func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB Disco y soy un %s", myPc.ram, myPc.disk, myPc.brand)
}

func Pointers() {
	a := 50
	b := &a

	functions.MyPrint("A value: ", a)
	functions.MyPrint("B pointer: ", b)
	functions.MyPrint("B pointer value: ", *b)

	functions.MyPrint("Set *b to 100")
	*b = 100
	functions.MyPrint("A value: ", a)

	functions.MyPrint("----------------------------")
	myPc := pc{ram: 32, disk: 5120, brand: "MSI"}
	myPc.ping()

	functions.MyPrint(myPc)
	myPc.duplicateRAM()

	functions.MyPrint(myPc)
	myPc.duplicateRAM()

	functions.MyPrint(myPc)
	myPc.duplicateRAM()

	functions.MyPrint(myPc)

}
