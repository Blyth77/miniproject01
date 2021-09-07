package main

import (
	"fmt"
)

// template - skal ændres på et tidspunkt
func fork() {
	var count int
	var state bool
}

func in_use() {
	state = true
	count++
}

func free() {
	state = false
}

func status() {
	fmt.Printf("%s, %s\n", count, state)
}
