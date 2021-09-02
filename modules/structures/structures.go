package structures

import "curso_basico_go/modules/functions"

type PublicCar struct {
	Brand string
	Year  int
}

type privateCar struct {
	brand string
	year  int
}

func MyStructures() {
	myCar := privateCar{brand: "Ford", year: 2020}
	functions.MyPrint(myCar)
	functions.MyPrint(myCar.brand)

	var otherCar privateCar
	otherCar.brand = "Ferrari"
	functions.MyPrint(otherCar)
}
