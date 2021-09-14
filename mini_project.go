package main

import (
	"fmt"
	"time"
)

func main() {
	// Fork1
	chan1, chan2, chan3, chan4 := make(chan int), make(chan int), make(chan int), make(chan int)

	// Fork2
	chan5, chan6, chan7, chan8 := make(chan int), make(chan int), make(chan int), make(chan int)

	//Fork 3
	chan9, chan10, chan11, chan12 := make(chan int), make(chan int), make(chan int), make(chan int)

	// Fork 4
	chan13, chan14, chan15, chan16 := make(chan int), make(chan int), make(chan int), make(chan int)

	// Fork 5
	chan17, chan18, chan19, chan20 := make(chan int), make(chan int), make(chan int), make(chan int)

	// Phil Input
	phil1, phil2, phil3, phil4, phil5 := make(chan int), make(chan int), make(chan int), make(chan int), make(chan int)

	go fork(chan1, chan2, chan3, chan4)
	go fork(chan5, chan6, chan7, chan8)
	go fork(chan9, chan10, chan11, chan12)
	go fork(chan13, chan14, chan15, chan16)
	go fork(chan17, chan18, chan19, chan20)

	go phil(chan20, chan19, chan2, chan1, phil1, "One")
	go phil(chan4, chan3, chan6, chan5, phil2, "Two")
	go phil(chan8, chan7, chan10, chan9, phil3, "Three")
	go phil(chan12, chan11, chan14, chan13, phil4, "Four")
	go phil(chan18, chan17, chan16, chan15, phil5, "Five") // Turned around

	fmt.Println("Dinners served!!")


	go readFromInput(phil1, phil2, phil3, phil4, phil5)

	for {

	}

}

func readFromInput(phil1, phil2, phil3, phil4, phil5 chan (int)) {
	var typeOf, idPhil, msg string

	fmt.Scan(&typeOf)
	fmt.Scan(&idPhil)
	fmt.Scan(&msg)

	if typeOf == "phil" {
		switch idPhil {
		case "1":
			phil1 <- 1
		case "2":
			phil2 <- 1
		case "3":
			phil3 <- 1
		case "4":
			phil4 <- 1
		case "5":
			phil5 <- 1
		default:
			break
		}
	}
}

func fork(chInLeft, chOutLeft, chInRight, chOutRight chan (int)) {
	timesUsed := 0

	for {
		// Modtager forespørgsel
		select {
			case <-chInLeft:
				{
					// Sender klarbesked
					chOutLeft <- 1

					// Modtager donebesked
					<-chInLeft
				}

			case <-chInRight:
				{
					// Sender klarbesked
					chOutRight <- 1

					// Modtager donebesked
					<-chInRight
				}
		}

		timesUsed++
	}
}

func phil(chInLeft, chOutLeft, chInRight, chOutRight, channelInput chan (int), name string) {
	timesEaten := 0
	status := "thinking"
	for {
		philMessages(channelInput, name, timesEaten, status)

		fmt.Printf("%s is thinking\n-----------------\n", name)
		time.Sleep(2 * time.Second)

		// Sends request
		chOutLeft <- 1

		// Recieve rdy
		<-chInLeft

		// Asks the other side
		chOutRight <- 1
		<-chInRight
		fmt.Printf("%s is eating\n", name)
		status = "eating"

		// Routine sleeps for 2 seconds
		time.Sleep(2 * time.Second)

		timesEaten++
		fmt.Printf("%s has eaten\n", name)

		// Sends "done"-msg
		chOutLeft <- 1
		chOutRight <- 1
		status = "thinking"
	}

}

func philMessages(channelOutput chan (int), name string, timesEaten int, status string) {
	select {
	case x := <-channelOutput:
		if x == 1 {
			fmt.Printf("Phil%s is %s", name, status)
		} else if x == 2 {
			fmt.Printf("Phil%s has eaten %s times!", name, timesEaten)
		}
	default:
	}
}
