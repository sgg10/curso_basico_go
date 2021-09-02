package maps

import "curso_basico_go/modules/functions"

func MyMaps() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Mariana"] = 15
	functions.MyPrint(m)

	for i, v := range m {
		functions.MyPrint(i, v)
	}

	// Find value
	value := m["Mariana"]
	functions.MyPrint(value)
	value_zero, ok := m["Pepito"]
	functions.MyPrint(value_zero, ok)
}
