package keywords

import "curso_basico_go/modules/functions"

func ContinueAndBreak(max_iter int) {
	for i := 0; i < max_iter; i++ {
		if i == 3 {
			continue
		}

		if i == int((max_iter/2)+1) {
			break
		}
		functions.MyPrint(i)
	}
}
