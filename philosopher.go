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
	p.right_fork.Lock()

	p.times_eaten++
	p.status = "eating"
	p.showStatus()

	time.Sleep(time.Second)

	p.left_fork.Unlock()
	p.right_fork.Unlock()

	p.status = "thinking"
	p.showStatus()
}

// id is..
func showStatus() {
	fmt.Printf("%s, %s\n", times_eaten, status)
}