package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	interrupted := make(chan os.Signal)
	// do anything
	fmt.Println("non-blocking, waiting to be interrupted")
	signal.Notify(interrupted, os.Interrupt)
	<-interrupted
	fmt.Println("interrupted, reached end")
}
