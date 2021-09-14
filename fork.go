package main

import (
	// "fmt"
)

func fork(chInLeft, chOutLeft, chInRight, chOutRight chan (int)) {
	timesUsed := 0

	for {
		// Receive request 
		select {
			case <-chInLeft:
				{
					chOutLeft <- 1 	// Recive READY-msg
					<-chInLeft // Sends DONE-msg
				}
			// Other side
			case <-chInRight:
				{
					chOutRight <- 1
					<-chInRight
				}
		}
		timesUsed++
	}
}