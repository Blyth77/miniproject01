package main

import (
	"fmt"
	"time"
)

/*
A PHILOSOPHER uses 2 channels to recive one fork, 1 channel for recieving commands and a name.
*/
func Philosopher(chInLeft, chOutLeft, chInRight, chOutRight, channelInput chan (int), channelOutput chan (string), name string) {
	timesEaten := 0
	status := "thinking"

	for {
		// CHECK if a command is incoming
		philMessages(channelInput, channelOutput, name, timesEaten, status)
		// THINKING:
		time.Sleep(2 * time.Second)
		//fmt.Printf("%s is thinking\n-----------------\n", name) // TEST

		// Try to EAT:
		takeFork(chInLeft, chOutLeft)
		takeFork(chInRight, chOutRight)

		// EATING:
		status = "eating"
		// fmt.Printf("%s is eating\n", name) // TEST
		time.Sleep(2 * time.Second) // Sleep
		timesEaten++
		putDownForks(chOutLeft, chOutRight) // Sends "done"-msg
		// CHECK if a command is INCOMING
		philMessages(channelInput, channelOutput, name, timesEaten, status)
		status = "thinking"
		// fmt.Printf("%s has eaten\n", name) // TEST
	}
}

func takeFork(forkIn, forkOut chan (int)) {
	forkOut <- 1 // Sends request
	<-forkIn     // Recieve rdy
}

func putDownForks(fork1, fork2 chan (int)) {
	fork1 <- 1
	fork2 <- 1
}

func philMessages(channelInput chan (int), channelOutput chan (string), name string, timesEaten int, status string) {
	select {
	case x := <-channelInput: // A msg IS incoming!
		if x == 1 {
			channelOutput <- fmt.Sprintf("PHILOSOPHER %s is %s\n PHILOSOPHER %s is not lstening anymore!\n", name, status, name)
			channelOutput <- fmt.Sprintf("PHILOSOPHER %s is not lstening anymore!\n", name)
		} else if x == 2 {
			channelOutput <- fmt.Sprintf("PHILOSOPHER %s has eaten %d time(s)!\n", name, timesEaten)
			channelOutput <- fmt.Sprintf("PHILOSOPHER %s is not listening anymore!\n", name)
		} else if x == 3 {
			channelOutput <- fmt.Sprintf("PHILOSOPHER %s is %s and has eaten %d time(s)!\n", name, status, timesEaten)
			channelOutput <- fmt.Sprintf("PHILOSOPHER %s is not listening anymore!\n", name)
		}
	default:
		// Stop blocking - if no msg is incoming
	}

}
