package main

import (
	"log"
	"os/exec"
	"sync"
)

func exe_cmd_wait(command string) {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go exe_cmd(command, wg)

	wg.Wait()
}

func exe_cmd(cmd string, wg *sync.WaitGroup) {
	log.Printf("Command: %s", cmd)

	out, err := exec.Command("/bin/sh", "-c", cmd).CombinedOutput()

	log.Printf("Output: %s", out)
	if err != nil {
		log.Fatalf("Error occured: %s", err)
	}

	wg.Done()
}
