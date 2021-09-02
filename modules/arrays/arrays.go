package arrays

import "curso_basico_go/modules/functions"

func Array() {
	var array [4]int
	array[0] = 1
	array[1] = 2
	functions.MyPrint(array, len(array), cap(array))
}

func Slice() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7}
	functions.MyPrint(slice, len(slice), cap(slice))

	// Slice methods
	functions.MyPrint(slice[0])
	functions.MyPrint(slice[:3])
	functions.MyPrint(slice[2:4])
	functions.MyPrint(slice[4:])

	// Append
	slice = append(slice, 8)
	functions.MyPrint(slice)

	new_slice := []int{9, 10}
	slice = append(slice, new_slice...)
	functions.MyPrint(slice)

	for i, v := range slice {
		functions.MyPrint(i, v)
	}
}
