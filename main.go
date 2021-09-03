package main

import (
	"curso_basico_go/modules/functions"
	//"curso_basico_go/modules/loops"
	//"curso_basico_go/modules/conditions"
	//"curso_basico_go/modules/switchCondition"
	//"curso_basico_go/modules/keywords"
	//"curso_basico_go/modules/arrays"
	//"curso_basico_go/modules/maps"
	//"curso_basico_go/modules/structures"
	//"curso_basico_go/modules/pointers"
	//"curso_basico_go/modules/interfaces"
	"curso_basico_go/modules/goroutines"
)

func main() {
	defer functions.MyPrint("Adios!")
	//functions.AllFunctions()
	//loops.ForLoops()
	//conditions.IfCondition()
	//switchCondition.Switch_condition()
	//keywords.ContinueAndBreak(13)
	//arrays.Array()
	//arrays.Slice()
	//maps.MyMaps()
	// car := structures.PublicCar{Brand: "Ferrari", Year: 2020}
	// functions.MyPrint(car)
	// structures.MyStructures()
	//pointers.Pointers()
	//interfaces.MyInterfaces()
	goroutines.GoRoutines()
}
