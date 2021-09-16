package main

import (
	"fmt"
	"os"
)

func main() {
	openingMsg()
	helpMsg()

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
		default:
			fmt.Println("Not understood? Please try again!")
		}
		if started {
			break
		}
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
	philIn1, philIn2, philIn3, philIn4, philIn5 := make(chan int, 10), make(chan int, 10), make(chan int, 10), make(chan int, 10), make(chan int, 10)
	// Phil output
	philOut1, philOut2, philOut3, philOut4, philOut5 := make(chan string, 10), make(chan string, 10), make(chan string, 10), make(chan string, 10), make(chan string, 10)
	// Fork's COMMAND-INPUT channels.
	forkIn1, forkIn2, forkIn3, forkIn4, forkIn5 := make(chan int, 10), make(chan int, 10), make(chan int, 10), make(chan int, 10), make(chan int, 10)

	// Fork output
	forkOut1, forkOut2, forkOut3, forkOut4, forkOut5 := make(chan string, 10), make(chan string, 10), make(chan string, 10), make(chan string, 10), make(chan string, 10)

	// Init all ROUTINES
	go Fork(chan1, chan2, chan3, chan4, forkIn1, forkOut1, "A")
	go Fork(chan5, chan6, chan7, chan8, forkIn2, forkOut2, "B")
	go Fork(chan9, chan10, chan11, chan12, forkIn3, forkOut3, "C")
	go Fork(chan13, chan14, chan15, chan16, forkIn4, forkOut4, "D")
	go Fork(chan17, chan18, chan19, chan20, forkIn5, forkOut5, "E")
	go Philosopher(chan20, chan19, chan2, chan1, philIn1, philOut1, "One")
	go Philosopher(chan4, chan3, chan6, chan5, philIn2, philOut2, "Two")
	go Philosopher(chan8, chan7, chan10, chan9, philIn3, philOut3, "Three")
	go Philosopher(chan12, chan11, chan14, chan13, philIn4, philOut4, "Four")
	go Philosopher(chan18, chan17, chan16, chan15, philIn5, philOut5, "Five") // Turned around - to avoid a DeadLocks

	// Dinners STARTING MSG -- program is now runnning
	fmt.Println("DINNERS SERVED!!")
	fmt.Println("------------------------------------------------")

	// Waits for and READS-USER-INPUT
	go readFromInput(philIn1, philIn2, philIn3, philIn4, philIn5, forkIn1, forkIn2, forkIn3, forkIn4, forkIn5)
	go printToTerminal(philOut1, philOut2, philOut3, philOut4, philOut5, forkOut1, forkOut2, forkOut3, forkOut4, forkOut5)
	for {
		// Program runs forever, until stopped.
	}

}

func readFromInput(phil1, phil2, phil3, phil4, phil5, fork1, fork2, fork3, fork4, fork5 chan (int)) {
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
		case "a":
			forkMsg("A", fork1)
		case "b":
			forkMsg("B", fork2)
		case "c":
			forkMsg("C", fork3)
		case "d":
			forkMsg("D", fork4)
		case "e":
			forkMsg("E", fork5)
		case "q":
			exit() // Program exit
		default:
			fmt.Println("Command not understood. Please try again!")
		}
	}
}

func printToTerminal(phil1, phil2, phil3, phil4, phil5, fork1, fork2, fork3, fork4, fork5 chan (string)) {
	for {
		var stringToPrint string
		select {
		case <-phil1:
			{
				stringToPrint = <-phil1
			}
		case <-phil2:
			{
				stringToPrint = <-phil2
			}
		case <-phil3:
			{
				stringToPrint = <-phil3
			}
		case <-phil3:
			{
				stringToPrint = <-phil4
			}
		case <-phil4:
			{
				stringToPrint = <-phil4
			}
		case <-phil5:
			{
				stringToPrint = <-phil5
			}
		case <-fork1:
			{
				stringToPrint = <-fork1
			}
		case <-fork2:
			{
				stringToPrint = <-fork2
			}
		case <-fork3:
			{
				stringToPrint = <-fork3
			}
		case <-fork4:
			{
				stringToPrint = <-fork4
			}
		case <-fork5:
			{
				stringToPrint = <-fork5
			}
		}
		fmt.Println(stringToPrint)
	}
}

func philMsg(name string, channel chan (int)) {
	// Show user options for interaction with a phil
	fmt.Printf("PHILOSOPHER %s is listening!\n", name)
	fmt.Println("Type a command to ask the philosopher something:")
	fmt.Println(" - type 's' to ask his status.")
	fmt.Println(" - type 'e' to ask how many times he has eaten.")
	fmt.Println(" - type 'z' for all info.")

	// WAIT for and RESPOND to phils command
	var msgSendSucces bool

	for {
		var command string
		fmt.Scan(&command)

		switch command {
		case "s":
			channel <- 1 // Status
			msgSendSucces = true
		case "e":
			channel <- 2 // TimesEaten
			msgSendSucces = true
		case "z":
			channel <- 3 // All
			msgSendSucces = true
		case "q":
			exit()
		default:
			fmt.Println("Command not understood. Please try again!")
		}
		if msgSendSucces {
			break
		}
	}
}

func forkMsg(id string, channel chan (int)) {
	// Show user options for interaction with a phil
	fmt.Printf("FORK %s is listening!\n", id)
	fmt.Println("Type a command to ask the fork something:")
	fmt.Println(" - type 's' to ask its status.")
	fmt.Println(" - type 'e' to ask how many times it has used.")
	fmt.Println(" - type 'z' for all info.")

	// WAIT for and RESPOND to phils command
	var msgSendSucces bool

	for {
		var command string
		fmt.Scan(&command)

		switch command {
		case "s":
			channel <- 1 // Status
			msgSendSucces = true
		case "e":
			channel <- 2 // TimesUsed
			msgSendSucces = true
		case "z":
			channel <- 3 // All
			msgSendSucces = true
		case "q":
			exit()
		default:
			fmt.Println("Command not understood. Please try again!")
		}
		if msgSendSucces {
			break
		}
	}
}

func openingMsg() {
	fmt.Printf("Welcome to the PHILOSOPHERS DINNER! :D :D :D\n\n")
	fmt.Println("To START the \"dinner\" press the 's'-key!")
	fmt.Printf("To ENDs the \"dinner\" press the 'q'-key!\n\n")
	fmt.Println("If you want to ask a philosopher something, press the corresponding key!: ")
	fmt.Println(" '1' ---  PHILOSOPHER ONE")
	fmt.Println(" '2' ---  PHILOSOPHER TWO")
	fmt.Println(" '3' ---  PHILOSOPHER THREE")
	fmt.Println(" '4' ---  PHILOSOPHER FOUR")
	fmt.Println(" '5' ---  PHILOSOPHER FIVE")
	fmt.Println("------------------------------------------------")
	fmt.Println("If you want to ask a fork something, press the corresponding key!: ")
	fmt.Println(" 'a' ---  FORK A")
	fmt.Println(" 'b' ---  FORK B")
	fmt.Println(" 'c' ---  FORK C")
	fmt.Println(" 'd' ---  FORK D")
	fmt.Println(" 'e' ---  FORK E")

	fmt.Println("If you want to ask a fork something, press the corresponding key!: ")
	fmt.Println("NOTE: all commands are followed by 'ENTER'!")
}

func helpMsg() {
	fmt.Print("Start program: s \n \t - followed by enter\n")
	fmt.Print("Quit program: q \n \t- followed by enter\n \t ")
	fmt.Print("While program is running: \n Call upon a philosopher, enter no. from 1-5 \n \t - followed by enter \n")
	fmt.Print("While philosopher is called upon: \n Get philosopher status: s \n \t - followed by enter \n")
	fmt.Print("Get number of times philosopher has eaten: e \n \t - followed by enter \n")
	fmt.Print("Get number of times the philosopher has eaten and his status: z \n \t")
}

func exit() {
	fmt.Printf("\nDinner has ended!!!\n")
	os.Exit(1)
}
