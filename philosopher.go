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
		// CHECK if a query is INCOMING
		philMessages(channelInput, channelOutput, name, timesEaten, status)
		// THINKING:
		time.Sleep(2 * time.Second)

		//fmt.Printf("%s is thinking\n", name) // TEST

		// Try to EAT:
		takeFork(chInLeft, chOutLeft)
		takeFork(chInRight, chOutRight)

		// EATING:
		philMessages(channelInput, channelOutput, name, timesEaten, status)
		status = "eating"
		philMessages(channelInput, channelOutput, name, timesEaten, status)

		//fmt.Printf("%s is eating\n", name) // TEST
		time.Sleep(2 * time.Second) // Sleep
		timesEaten++
		putDownForks(chOutLeft, chOutRight) // Sends "done"-msg
		// CHECK if a query is INCOMING
		philMessages(channelInput, channelOutput, name, timesEaten, status)
		status = "thinking"
		//fmt.Printf("%s has eaten\n", name) // TEST
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
			channelOutput <- "" // Signals select-block to expect a msg on the channel.
			channelOutput <- fmt.Sprintf("PHILOSOPHER %s is %s\nPHILOSOPHER %s is not listening anymore!\n", name, status, name)
		} else if x == 2 {
			channelOutput <- ""
			channelOutput <- fmt.Sprintf("PHILOSOPHER %s has eaten %d time(s)!\nPHILOSOPHER %s is not listening anymore!\n", name, timesEaten, name)
		} else if x == 3 {
			channelOutput <- ""
			channelOutput <- fmt.Sprintf("PHILOSOPHER %s is %s and has eaten %d time(s)!\nPHILOSOPHER %s is not listening anymore!\n", name, status, timesEaten, name)
		} else if x == 4 {
			channelOutput <- ""
			channelOutput <- fmt.Sprintf("PHILOSOPHER %s is %s\n", name, status)
		} else if x == 5 {
			channelOutput <- ""
			channelOutput <- fmt.Sprintf("PHILOSOPHER %s has eaten %d time(s)!\n", name, timesEaten)
		} else if x == 6 {
			channelOutput <- ""
			channelOutput <- fmt.Sprintf("PHILOSOPHER %s is %s and has eaten %d time(s)!\n", name, status, timesEaten)
		}
	default:
		// Stop blocking - if no msg is incoming
	}

}
