package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal) // channel for OS signals

	// when the SIGINT signal is triggered, add the signal to the sigs channel
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go goFunc("A")
	go goFunc("B")
	go goFunc("C")

	<-sigs // block until we receive a OS signal.
}

func goFunc(str string) {
	for {
		fmt.Println(str + " started")
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(str + " finished")
	}
}
