package main

import (
	"log"
	"os/exec"
	"sync"
	"time"
)

func executeCommandAndWait(command string, timeout time.Duration) {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go executeCommand(command, wg)

	if waitTimeout(wg, timeout) {
		log.Fatalf("Command timed out: %s", command)
	}
}

func executeCommand(cmd string, wg *sync.WaitGroup) {
	log.Printf("Command: %s", cmd)

	out, err := exec.Command("/bin/sh", "-c", cmd).CombinedOutput()

	log.Printf("Output: %s", out)
	if err != nil {
		log.Fatalf("Error occured: %s", err)
	}

	wg.Done()
}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}
