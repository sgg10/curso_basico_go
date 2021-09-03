package chanels

import "curso_basico_go/modules/functions"

func say(text string, c chan<- string) {
	c <- text
}

func message(text string, c chan<- string) {
	c <- text
}

func Chanels() {
	c := make(chan string, 1)

	functions.MyPrint("Hello")

	go say("Bye", c)

	functions.MyPrint(<-c)

	functions.MyPrint("-------------------------------")
	c2 := make(chan string, 2)

	c2 <- "message 1"
	c2 <- "message 2"

	functions.MyPrint(len(c2), cap(c2))

	close(c2) // Channel closed, you cannot send anymore

	for message := range c2 {
		functions.MyPrint(message)
	}

	// Select
	emails := make(chan string)
	emails2 := make(chan string)

	go message("message3", emails)
	go message("message4", emails2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-emails:
			functions.MyPrint("Email retrieved from emails", m1)
		case m2 := <-emails2:
			functions.MyPrint("Email retrieved from emails2", m2)
		}
	}

}
