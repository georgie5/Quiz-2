package main

import (
	"fmt"
	"time"
)

// Go routines
func dosmthing() {
	fmt.Println("does smthing")
}
func becreative() {
	fmt.Println("being creative")
}

// a main is a go routine
func main() {
	go dosmthing()
	go becreative()
	//main to wait
	time.Sleep(time.Second)
}
