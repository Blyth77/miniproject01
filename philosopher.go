package main

import (
	"fmt"
	"time"
)

type Phil struct{
	var id int
	var times_eaten int
	var status string
	var left_fork, right_fork *Fork
}

func eating(p Phil) {
	p.left_fork.Lock()
	p.showForkPickUp(left_fork)

	p.right_fork.Lock()
	p.showForkPickUp(right_fork)

	p.times_eaten++
	p.status = "eating"
	p.showStatus()

	time.Sleep(time.Second)

	p.left_fork.Unlock()
	p.right_fork.Unlock()

	p.status = "thinking"
	p.showStatus()
}

func showForkPickUp(fork Fork){
	fmt.Printf("%s, %s\n", "Picked up: ", fork)
}

// id is..
func showStatus() {
	fmt.Printf("%s, %s\n", times_eaten, status)
}