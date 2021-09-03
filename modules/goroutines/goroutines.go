package goroutines

import (
	"curso_basico_go/modules/functions"
	"sync"
	"time"
)

func say(message string, wg *sync.WaitGroup) {
	defer wg.Done()
	functions.MyPrint(message)
	time.Sleep(time.Second * 1)
	functions.MyPrint("Againg: ", message)
}

func GoRoutines() {
	var wg sync.WaitGroup

	functions.MyPrint("Hello")
	wg.Add(1)

	go say("World!", &wg)

	wg.Add(1)
	go func(message string, wg *sync.WaitGroup) {
		defer wg.Done()
		functions.MyPrint(message)
	}("Hi!", &wg)

	wg.Wait()
}
