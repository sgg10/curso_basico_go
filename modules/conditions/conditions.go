package conditions

import (
	"curso_basico_go/modules/functions"
)

func IfCondition() {
	v1 := 1
	v2 := 2

	if v1 == 1 && v2 == 2 {
		functions.MyPrint("Es 1 y 2")
	} else {
		functions.MyPrint("No es 1 ni 2")
	}
}
