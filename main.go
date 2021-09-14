package main

import (
	"fmt"
	"os"
)

func main() {
	openingMsg()
	
	// WAIT for PROGRAM-START or exit
	for {
		var started bool
		var inputKey string
		fmt.Scan(&inputKey)
		switch inputKey {
			case "s":
				started = true
			case "q":
				exit()
			default: fmt.Println("Not understood? Please try again!")
		}
		if started { break }
	}


	// Init all CHANNELS

	// Fork 1
	chan1, chan2, chan3, chan4 := make(chan int), make(chan int), make(chan int), make(chan int)
	// Fork 2
	chan5, chan6, chan7, chan8 := make(chan int), make(chan int), make(chan int), make(chan int)
	// Fork 3
	chan9, chan10, chan11, chan12 := make(chan int), make(chan int), make(chan int), make(chan int)
	// Fork 4
	chan13, chan14, chan15, chan16 := make(chan int), make(chan int), make(chan int), make(chan int)
	// Fork 5
	chan17, chan18, chan19, chan20 := make(chan int), make(chan int), make(chan int), make(chan int)


	// Phil's COMMAND-INPUT channels. 
	phil1, phil2, phil3, phil4, phil5 := make(chan int), make(chan int), make(chan int), make(chan int), make(chan int)


	// Init all ROUTINES
	go fork(chan1, chan2, chan3, chan4)
	go fork(chan5, chan6, chan7, chan8)
	go fork(chan9, chan10, chan11, chan12)
	go fork(chan13, chan14, chan15, chan16)
	go fork(chan17, chan18, chan19, chan20)
	go philosopher(chan20, chan19, chan2, chan1, phil1, "One")
	go philosopher(chan4, chan3, chan6, chan5, phil2, "Two")
	go philosopher(chan8, chan7, chan10, chan9, phil3, "Three")
	go philosopher(chan12, chan11, chan14, chan13, phil4, "Four")
	go philosopher(chan18, chan17, chan16, chan15, phil5, "Five") // Turned around - to avoid a DeadLocks

	// Dinners STARTING MSG -- program is now runnning
	fmt.Println("DINNERS SERVED!!")
	fmt.Println("------------------------------------------------")

	// Waits for and READS-USER-INPUT
	go readFromInput(phil1, phil2, phil3, phil4, phil5)

	for {
		// Program runs forever, until stopped.
	}

}

func readFromInput(phil1, phil2, phil3, phil4, phil5 chan (int)) {
	for {
		var input string
		fmt.Scan(&input)

		// Choose which PHIL-TO-CONTACT
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
				exit() // Program exit
			default:
				fmt.Println("Command not understood. Please try again!")
				break 
		}
	}
}

func philMsg(name string, channel chan (int)) {
	// Show user options for interaction with a phil
	fmt.Printf("PHILOSOPHER %s is listening!\n", name)
	fmt.Println("Type 'e' to ask how many times he has eaten.")
	fmt.Println(" -or type 's' to ask his status.")

	// WAIT for and RESPOND to phils command
	var command string
	fmt.Scan(&command)

	switch command {
		case "s":
			channel <- 1 // Status
		case "e":
			channel <- 2 // TimesEaten
		case "q":
			exit()
	}
}

func openingMsg() {
	fmt.Printf("Welcome to the PHILOSOPHERS DINNER! :D :D :D\n\n")
	fmt.Println("To START the \"dinner\" press the 's'-key!")
	fmt.Printf("To ENDs the \"dinner\" press the 'q'-key!\n\n")
	fmt.Println("If you want to ask a philosopher something, press the corresponding key!: " )
	fmt.Println(" '1' ---  PHILOSOPHER ONE")
	fmt.Println(" '2' ---  PHILOSOPHER TWO")
	fmt.Println(" '3' ---  PHILOSOPHER THREE")
	fmt.Println(" '4' ---  PHILOSOPHER FOUR")
	fmt.Println(" '5' ---  PHILOSOPHER FIVE")
	fmt.Println("NOTE: all commands are followed by 'ENTER'!")
}

func exit() {
	fmt.Printf("\nDinner has ended!!!\n")
	os.Exit(1)
}