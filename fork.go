package main

import (
	"fmt"
)

func Fork(chInLeft, chOutLeft, chInRight, chOutRight, channelInput chan (int), id string) {
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
			fmt.Printf("Fork %s is %s\n", id, status)
		} else if x == 2 {
			fmt.Printf("Fork %s has eaten %d time(s)!\n", id, timesUsed)
		}
	default:
		// Stop blocking - if no msg is incoming
	}
}
