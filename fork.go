package main

import (
	"fmt"
	"sync"
)

// template - skal ændres på et tidspunkt

type Fork struct{
	var times_used int
	var in_use bool
	sync.Mutex
}


func takeFork() {
	state = true
	counter++
}

func status() bool {
	return status
}

func showStatus() {
	fmt.Printf("%s, %s\n", counter, state)
}
