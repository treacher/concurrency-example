/*
	Concurrency Example:

	This example executes three go functions and when the application receives a SIGINT it makes
	sure all of the functions have finished executing before the application exits.

	To prove this we pipe the output to a file and make sure we have the same amount starts and finishes.
*/

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func safeGoFunc(wg *sync.WaitGroup, stop chan struct{}, str string) {
	defer wg.Done() // when this function finishes, notify the wait group that this function has finished.

	wg.Add(1) // notify the wait group that we're executing a function.

	for {
		select {
		case <-stop:
			return
		default:
			fmt.Println(str + " started")
			time.Sleep(1000 * time.Millisecond)
			fmt.Println(str + " finished")
		}
	}
}

func main() {
	sigs := make(chan os.Signal) // channel for OS signals
	stop := make(chan struct{})  // channel for signalling to go functions to stop running

	// when the SIGINT signal is triggered, add the signal to the sigs channel
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	// Wait group for keeping track of how many go functions are currently running.
	wg := &sync.WaitGroup{}
	defer wg.Wait() // when this function finishes wait until all of the go routines have finished

	go safeGoFunc(wg, stop, "A")
	go safeGoFunc(wg, stop, "B")
	go safeGoFunc(wg, stop, "C")

	<-sigs // block until we receive a OS signal.

	close(stop) // close the stop channel to tell the go routines to finish up
}
