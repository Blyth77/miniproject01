package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
	openingMsg()
	
	for {
		var inputKey string
		fmt.Scan(&inputKey)
		if inputKey == "s" {
			break
		} else {
			fmt.Println("Not understood? Please try again!")
		}

	}





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

func openingMsg() {
	fmt.Println("Welcome to the PHILOSOPHERS DINNER! :D :D :D\n")
	fmt.Println("To START the \"dinner\" press the 's'-key!")
	fmt.Println("To ENDs the \"dinner\" press the 'q'-key!\n")
	fmt.Println("If you want to ask a philosopher something, press the corresponding key!: " )
	fmt.Println(" '1' ---  PHILOSOPHER ONE")
	fmt.Println(" '2' ---  PHILOSOPHER TWO")
	fmt.Println(" '3' ---  PHILOSOPHER THREE")
	fmt.Println(" '4' ---  PHILOSOPHER FOUR")
	fmt.Println(" '5' ---  PHILOSOPHER FIVE")
	fmt.Println("NOTE: all commands are followed by 'ENTER'!")
}

func readFromInput(phil1, phil2, phil3, phil4, phil5 chan (int)) {
	for {
		var input string

		fmt.Scan(&input)

		switch input {
			case "1":
				philMsg("ONE", phil1)
			case "2":
				philMsg("TWO", phil2)
			case "3":
				philMsg("THREE", phil3)
			case "4":
				philMsg("FOUR", phil4)
			case "5":
				philMsg("FIVE", phil5)
			case "q":
				exit()
			default:
				break
		}
	}
}

func exit() {
	fmt.Printf("\nDinner has ended!!!\n")
	os.Exit(1)
}

func philMsg(name string, channel chan (int)) {
	fmt.Printf("PHILOSOPHER %s is listening!\n", name)
	fmt.Println("Type 'e' to ask how many times he has eaten.")
	fmt.Println(" -or type 's' to ask his status.")


	var command string
	fmt.Scan(&command)
	switch command {
	case "s":
		channel <- 1
	case "e":
		channel <- 2
	case "q":
		exit()
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

		//fmt.Printf("%s is thinking\n-----------------\n", name)
		time.Sleep(2 * time.Second)

		// Sends request
		chOutLeft <- 1

		// Recieve rdy
		<-chInLeft

		// Asks the other side
		chOutRight <- 1
		<-chInRight
		// fmt.Printf("%s is eating\n", name)
		status = "eating"
		philMessages(channelInput, name, timesEaten, status)

		// Routine sleeps for 2 seconds
		time.Sleep(2 * time.Second)

		timesEaten++
		// fmt.Printf("%s has eaten\n", name)

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
			fmt.Printf("Phil %s is %s\n", name, status)
		} else if x == 2 {
			fmt.Printf("Phil%s has eaten %d time(s)!\n", name, timesEaten)
		}
	default:
	}
}
