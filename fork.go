package main

import (
	"fmt"
)

func Fork(chInLeft, chOutLeft, chInRight, chOutRight, channelInput chan (int), channelOutput chan (string), id string) {
	timesUsed := 0
	status := "free"

	for {
		// CHECK if a command is INCOMING
		forkMessages(channelInput, id, timesUsed, status)

		// Receive request
		select {
		case <-chInLeft:
			{
				chOutLeft <- 1 // Recive READY-msg
				status = "in use"
				<-chInLeft // Sends DONE-msg
			}
		// Other side
		case <-chInRight:
			{
				chOutRight <- 1
				status = "in use"
				<-chInRight
			}
		}
		// CHECK if a command is INCOMING
		forkMessages(channelInput, id, timesUsed, status)
		status = "free"
		timesUsed++
	}
}

func forkMessages(channelOutput chan (int), id string, timesUsed int, status string) {
	select {
	case x := <-channelOutput: // A msg IS incoming!
		if x == 1 {
			fmt.Printf("FORK %s is %s\n", id, status)
			fmt.Printf("FORK %s is not listening anymore!\n", id)
		} else if x == 2 {
			fmt.Printf("FORK %s has been used %d time(s)!\n", id, timesUsed)
			fmt.Printf("FORK %s is not listening anymore!\n", id)
		} else if x == 3 {
			fmt.Printf("FORK %s is %s and has been used %d time(s)!\n", id, status, timesUsed)
			fmt.Printf("FORK %s is not listening anymore!\n", id)
		}
	default:
		// Stop blocking - if no msg is incoming
	}
}
