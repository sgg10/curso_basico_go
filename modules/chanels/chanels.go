package chanels

import "curso_basico_go/modules/functions"

func say(text string, c chan<- string) {
	c <- text
}

func Chanels() {
	c := make(chan string, 1)

	functions.MyPrint("Hello")

	go say("Bye", c)

	functions.MyPrint(<-c)
}
