package switchCondition

import "curso_basico_go/modules/functions"

func Switch_condition() {
	switch modulo := 4 % 2; modulo {
	case 0:
		functions.MyPrint("Es par")
	default:
		functions.MyPrint("No es par")
	}

	//Without condition
	value := 200
	switch {
	case value > 100:
		functions.MyPrint("Es mayor a 100")
	case value < 0:
		functions.MyPrint("Es menor a 0")
	default:
		functions.MyPrint("No condition")
	}
}
