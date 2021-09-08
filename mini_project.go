package main

import (
	"fmt"
)

func main() {
	number_of := 5

	phils := [number_of]Phil
	forks := [number_of]Fork

	for i := 0; i < number_of; i++ {
		forks[i] = new(Fork)

	}

	for i := 0; i < number_of; i++ {
		phils[i] = new(Phil{
			id : i+1, status : "thinking", left_fork: forks[i], right_fork: forks[(i+1)%number_of]
		})
		go eating(phils[i])
	}


	// DECLARE:
	// Array med philosopher 

	// Array med forks

	// Starte programmet 

}
