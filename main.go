package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Enter in main")

	go callGo()
	time.Sleep(1 * time.Second)

	// using chanels
	chanVal := make(chan bool)

	go callGoChan(chanVal)
	<-chanVal

	fmt.Println("Executed main...")

}

// callGo :
func callGo() {
	fmt.Println("Calling callGo...")
}

// callGoChan :
func callGoChan(chanVal chan bool) {
	fmt.Println("Calling callGoChan...")

	chanVal <- true
}
