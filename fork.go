package main

import (
	"fmt"
)

func Fork(chInLeft, chOutLeft, chInRight, chOutRight, channelInput chan (int), channelOutput chan (string), id string) {
	timesUsed := 0
	status := "free"

	for {
		// CHECK if a query is INCOMING
		forkMsg(channelInput, channelOutput, id, timesUsed, status)

		// Receive request
		select {
		case <-chInLeft:
			{
				chOutLeft <- 1 // Send READY-msg
				status = "in use"
				<-chInLeft // Recieve DONE-msg
			}
		// Other side
		case <-chInRight:
			{
				chOutRight <- 1
				status = "in use"
				<-chInRight
			}
		}
		// CHECK if a query is INCOMING
		forkMsg(channelInput, channelOutput, id, timesUsed, status)
		status = "free"
		timesUsed++
	}
}

func forkMsg(channelInput chan (int), channelOutput chan (string), id string, timesUsed int, status string) {
	select {
	case x := <-channelInput: // A msg IS incoming!
		if x == 1 {
			channelOutput <- "" // Signals select-block to expect a msg on the channel.
			channelOutput <- fmt.Sprintf("FORK %s is %s\nFORK %s is not listening anymore!\n", id, status, id)
		} else if x == 2 {
			channelOutput <- ""
			channelOutput <- fmt.Sprintf("FORK %s has been used %d time(s)!\nFORK %s is not listening anymore!\n", id, timesUsed, id)
		} else if x == 3 {
			channelOutput <- ""
			channelOutput <- fmt.Sprintf("FORK %s is %s and has been used %d time(s)!\nFORK %s is not listening anymore!\n", id, status, timesUsed, id)
		} else if x == 4 {
			channelOutput <- ""
			channelOutput <- fmt.Sprintf("FORK %s is %s\n", id, status)
		} else if x == 5 {
			channelOutput <- ""
			channelOutput <- fmt.Sprintf("FORK %s has been used %d time(s)!\n", id, timesUsed)
		} else if x == 6 {
			channelOutput <- ""
			channelOutput <- fmt.Sprintf("FORK %s is %s and has been used %d time(s)!\n", id, status, timesUsed)
		}
	default:
		// Stop blocking - if no msg is incoming
	}
}
