package main

import (
	"fmt"
	"time"
)

/* 
A PHILOSOPHER uses 2 channels to recive one fork, 1 channel for recieving commands and a name.
*/
func Philosopher(chInLeft, chOutLeft, chInRight, chOutRight, channelInput chan (int), name string) {
	timesEaten := 0
	status := "thinking"

	for {
		// CHECK if a command is incoming
		philMessages(channelInput, name, timesEaten, status)
		
		// THINKING:
		time.Sleep(2 * time.Second) 
		//fmt.Printf("%s is thinking\n-----------------\n", name) // TEST


		// Try to EAT:
		takeFork(chInLeft, chOutLeft)
		takeFork(chInRight, chOutRight)

		// CHECK if a command is INCOMING

		// EATING:
		status = "eating"
		// fmt.Printf("%s is eating\n", name) // TEST
		time.Sleep(2 * time.Second) // Sleep
		timesEaten++
		putDownForks(chOutLeft, chOutRight) // Sends "done"-msg
		philMessages(channelInput, name, timesEaten, status)
		status = "thinking"
		// fmt.Printf("%s has eaten\n", name) // TEST
	}
}

func takeFork(forkIn, forkOut chan(int)) {
	forkOut <- 1 // Sends request
	<-forkIn 	// Recieve rdy
}

func putDownForks(fork1, fork2 chan(int)) {
	fork1 <- 1
	fork2 <- 1
}

func philMessages(channelOutput chan (int), name string, timesEaten int, status string) {
	select {
		case x := <-channelOutput: // A msg IS incoming!
			if x == 1 {
				fmt.Printf("PHILOSOPHER %s is %s\n", name, status)
				fmt.Printf("PHILOSOPHER %s is not lstening anymore!\n", name)
			} else if x == 2 {
				fmt.Printf("PHILOSOPHER %s has eaten %d time(s)!\n", name, timesEaten)
				fmt.Printf("PHILOSOPHER %s is not listening anymore!\n", name)
			} else if x == 3 {
				fmt.Printf("PHILOSOPHER %s is %s and has eaten %d time(s)!\n", name, status, timesEaten)
				fmt.Printf("PHILOSOPHER %s is not listening anymore!\n", name)
			} 
		default:
			// Stop blocking - if no msg is incoming
	}

}