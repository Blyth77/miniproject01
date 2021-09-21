package main

import (
	"fmt"
	"os"
)

func getCount() int {
	return 5
}

func main() {
	openingMsg()

	// WAIT for PROGRAM-START or exit
	startMenu(getUserInput)

	// Init all CHANNELS

	// Fork A connections
	chan1, chan2, chan3, chan4 := make(chan int), make(chan int), make(chan int), make(chan int)
	// Fork B connections
	chan5, chan6, chan7, chan8 := make(chan int), make(chan int), make(chan int), make(chan int)
	// Fork C connections
	chan9, chan10, chan11, chan12 := make(chan int), make(chan int), make(chan int), make(chan int)
	// Fork D connections
	chan13, chan14, chan15, chan16 := make(chan int), make(chan int), make(chan int), make(chan int)
	// Fork E connections
	chan17, chan18, chan19, chan20 := make(chan int), make(chan int), make(chan int), make(chan int)

	// Phil's Query-INPUT channels.
	philIn1, philIn2, philIn3, philIn4, philIn5 := make(chan int, 10), make(chan int, 10), make(chan int, 10), make(chan int, 10), make(chan int, 10)
	// Phil OUTPUT
	philOut1, philOut2, philOut3, philOut4, philOut5 := make(chan string, 10), make(chan string, 10), make(chan string, 10), make(chan string, 10), make(chan string, 10)
	// Fork's Query-INPUT channels.
	forkIn1, forkIn2, forkIn3, forkIn4, forkIn5 := make(chan int, 10), make(chan int, 10), make(chan int, 10), make(chan int, 10), make(chan int, 10)
	// Fork OUTPUT
	forkOut1, forkOut2, forkOut3, forkOut4, forkOut5 := make(chan string, 10), make(chan string, 10), make(chan string, 10), make(chan string, 10), make(chan string, 10)

	// Init all ROUTINES
	go Fork(chan1, chan2, chan3, chan4, forkIn1, forkOut1, "A")
	go Fork(chan5, chan6, chan7, chan8, forkIn2, forkOut2, "B")
	go Fork(chan9, chan10, chan11, chan12, forkIn3, forkOut3, "C")
	go Fork(chan13, chan14, chan15, chan16, forkIn4, forkOut4, "D")
	go Fork(chan17, chan18, chan19, chan20, forkIn5, forkOut5, "E")
	go Philosopher(chan20, chan19, chan2, chan1, philIn1, philOut1, "ONE")
	go Philosopher(chan4, chan3, chan6, chan5, philIn2, philOut2, "TWO")
	go Philosopher(chan8, chan7, chan10, chan9, philIn3, philOut3, "THREE")
	go Philosopher(chan12, chan11, chan14, chan13, philIn4, philOut4, "FOUR")
	go Philosopher(chan18, chan17, chan16, chan15, philIn5, philOut5, "FIVE") // Turned around - to avoid a DeadLock

	// Dinners STARTING MSG -- program is now runnning
	startMsg()

	// Waits for and READS-USER-INPUT
	go queryEntityFromInput(getUserInput, philIn1, philIn2, philIn3, philIn4, philIn5, forkIn1, forkIn2, forkIn3, forkIn4, forkIn5)

	// Sends OUTPUT-TO-TERMINAL
	go outputFromUserQueries(output, philOut1, philOut2, philOut3, philOut4, philOut5, forkOut1, forkOut2, forkOut3, forkOut4, forkOut5)

	for {
		// Program runs forever, until stopped.
	}
}

func startMenu(input func() string) {
	for {
		var started bool
		switch input() {
		case "s":
			started = true
		case "q":
			exit()
		case "h":
			helpMsg()
		default:
			output("Query not understood. Please try again!")
		}
		if started {
			break
		}
	}
}

// IN from user
func getUserInput() string {
	var input string
	fmt.Scan(&input)
	return input
}

// OUT to user
func output(output string) {
	fmt.Println(output) // Outputs to terminal, but it can be changed!
}

func queryEntityFromInput(input func() string, phil1, phil2, phil3, phil4, phil5, fork1, fork2, fork3, fork4, fork5 chan (int)) {
	for {
		// Choose which PHIL-TO-CONTACT
		switch input() {
		case "1":
			query(input, philQueryOptions, "ONE", phil1)
		case "2":
			query(input, philQueryOptions, "TWO", phil2)
		case "3":
			query(input, philQueryOptions, "THREE", phil3)
		case "4":
			query(input, philQueryOptions, "FOUR", phil4)
		case "5":
			query(input, philQueryOptions, "FIVE", phil5)
		case "a":
			query(input, forkQueryOptions, "A", fork1)
		case "b":
			query(input, forkQueryOptions, "B", fork2)
		case "c":
			query(input, forkQueryOptions, "C", fork3)
		case "d":
			query(input, forkQueryOptions, "D", fork4)
		case "e":
			query(input, forkQueryOptions, "E", fork5)
		case "p":
			queryAllPhil(input, allPhilQueryOptions, phil1, phil2, phil3, phil4, phil5)
		case "f":
			queryAllForks(input, allForkQueryOptions, fork1, fork2, fork3, fork4, fork5)
		case "q":
			exit() // Program exit
		case "h":
			helpMsg()
		default:
			output("Query not understood. Please try again!")
		}
	}
}

func outputFromUserQueries(queryResponse func(string), phil1, phil2, phil3, phil4, phil5, fork1, fork2, fork3, fork4, fork5 chan (string)) {
	for {
		select {
		case <-phil1:
			queryResponse(<-phil1)
		case <-phil2:
			queryResponse(<-phil2)
		case <-phil3:
			queryResponse(<-phil3)
		case <-phil4:
			queryResponse(<-phil4)
		case <-phil5:
			queryResponse(<-phil5)
		case <-fork1:
			queryResponse(<-fork1)
		case <-fork2:
			queryResponse(<-fork2)
		case <-fork3:
			queryResponse(<-fork3)
		case <-fork4:
			queryResponse(<-fork4)
		case <-fork5:
			queryResponse(<-fork5)
		}
	}
}

func queryAllPhil(input func() string, QueryOptions func() string, chan1, chan2, chan3, chan4, chan5 chan (int)) {

	output(allPhilQueryOptions())

	var validQuery bool

	for {
		switch input() {
		case "s":
			chan1 <- 4
			chan2 <- 4
			chan3 <- 4
			chan4 <- 4
			chan5 <- 4
			validQuery = true
		case "e":
			chan1 <- 5
			chan2 <- 5
			chan3 <- 5
			chan4 <- 5
			chan5 <- 5
			validQuery = true
		case "z":
			chan1 <- 6
			chan2 <- 6
			chan3 <- 6
			chan4 <- 6
			chan5 <- 6
			validQuery = true
		case "q":
			exit()
		default:
			output("Query not understood. Please try again!")
		}
		if validQuery {
			break
		}
	}
}

func queryAllForks(input func() string, QueryOptions func() string, chan1, chan2, chan3, chan4, chan5 chan (int)) {

	output(allForkQueryOptions())

	var validQuery bool

	for {
		switch input() {
		case "s":
			chan1 <- 4
			chan2 <- 4
			chan3 <- 4
			chan4 <- 4
			chan5 <- 4
			validQuery = true
		case "e":
			chan1 <- 5
			chan2 <- 5
			chan3 <- 5
			chan4 <- 5
			chan5 <- 5
			validQuery = true
		case "z":
			chan1 <- 6
			chan2 <- 6
			chan3 <- 6
			chan4 <- 6
			chan5 <- 6
			validQuery = true
		case "q":
			exit()
		default:
			output("Query not understood. Please try again!")
		}
		if validQuery {
			break
		} // RESTART if not valid!
	}

}

func query(input func() string, queryOptions func(string) string, id string, channel chan (int)) {

	// Show USER OPTIONS for interaction with an entity
	output(queryOptions(id))

	var validQuery bool

	// WAIT for and RESPOND to entity query
	for {
		switch input() {
		case "s":
			channel <- 1 // Status
			validQuery = true
		case "e":
			channel <- 2 // Counter
			validQuery = true
		case "z":
			channel <- 3 // All info
			validQuery = true
		case "h":
			queryOptions(id)
		case "q":
			exit()
		default:
			output("Query not understood. Please try again!")
		}
		if validQuery {
			break
		} // RESTART if not valid!
	}
}

// INSTRUCTIONS AND INFO for terminal
func openingMsg() {
	output("------------------------------------------------\n" +
		"Welcome to the PHILOSOPHERS DINNER! :D :D :D\n" +
		"------------------------------------------------\n" +
		"To START the \"dinner\" press the 's'-key!\n" +
		"To ENDs the \"dinner\" press the 'q'-key!\n" +
		"Need help type 'h'!\n" +
		"------------------------------------------------\n" +
		"If you want to ask a philosopher something, press the corresponding key!: \n" +
		" '1' ---  PHILOSOPHER ONE\n" +
		" '2' ---  PHILOSOPHER TWO\n" +
		" '3' ---  PHILOSOPHER THREE\n" +
		" '4' ---  PHILOSOPHER FOUR\n" +
		" '5' ---  PHILOSOPHER FIVE\n" +
		" To contact all philosophers press 'p' \n" +
		"------------------------------------------------\n" +
		"If you want to ask a fork something, press the corresponding key!: \n" +
		" 'a' ---  FORK A\n" +
		" 'b' ---  FORK B\n" +
		" 'c' ---  FORK C\n" +
		" 'd' ---  FORK D\n" +
		" 'e' ---  FORK E\n" +
		" To contact all forks press 'f' \n" +
		"------------------------------------------------\n" +
		"NOTE: all queries are followed by 'ENTER'!\n")
}

func helpMsg() {
	output("------------------------------------------------\n" +
		"Start program: s \n \t - followed by enter\n" +
		"Quit program: q \n \t- followed by enter\n \t " +
		"While program is running: \nCall upon a philosopher, enter no. from 1-5 \n \t - followed by enter \n" +
		"While philosopher is called upon: \nGet philosopher status: s \n \t - followed by enter \n" +
		"Get number of times philosopher has eaten: e \n \t - followed by enter \n" +
		"Get number of times the philosopher has eaten and his status: z \n \t" +
		"Get attention of all philosopher: p \n \t - followed by enter\n" +
		"Get attention of all forks: f \n \t - followed by enter\n" +
		"------------------------------------------------\n")
}

func philQueryOptions(name string) string {
	return "------------------------------------------------\n" +
		fmt.Sprintf("PHILOSOPHER %s is listening!\n", name) +
		"Type a query to ask the philosopher something:\n" +
		" - type 's' to ask his status.\n" +
		" - type 'e' to ask how many times he has eaten.\n" +
		" - type 'z' for all info.\n" +
		"------------------------------------------------\n"
}

func forkQueryOptions(id string) string {
	return "------------------------------------------------\n" +
		fmt.Sprintf("FORK %s is listening!\n", id) +
		"Type a query to ask the fork something:\n" +
		" - type 's' to ask its status.\n" +
		" - type 'e' to ask how many times it has been used.\n" +
		" - type 'z' for all info.\n" +
		"------------------------------------------------\n"
}

func allPhilQueryOptions() string {
	return "------------------------------------------------\n" +
		fmt.Sprintf("All PHILOSOPHERS are listening\n") +
		"Type a query to ask the philosophers something:\n" +
		" - type 's' to ask their status." +
		" - type 'e' to ask how many times they have eaten.\n" +
		" - type 'z' for all info.\n" +
		"------------------------------------------------\n"
}

func allForkQueryOptions() string {
	return "------------------------------------------------\n" +
		fmt.Sprintf("All FORKS is listening!\n") +
		"Type a query to ask the forks something:\n" +
		" - type 's' to ask their status." +
		" - type 'e' to ask how many times they have been used.\n" +
		" - type 'z' for all info.\n" +
		"------------------------------------------------\n"
}

func startMsg() {
	output("------------------------------------------------\n" +
		"	DINNERS IS SERVED!!\n" +
		"------------------------------------------------\n")
}

func exit() {
	output("------------------------------------------------\n" +
		"	Dinner has ended!!!\n" +
		"------------------------------------------------\n")
	os.Exit(0)
}
