package main

import (
	"log"
	"os/exec"
	"sync"
)

func exe_cmd_wait(command string) {
	wg := new(sync.WaitGroup)
	commands := []string{command}
	for _, str := range commands {
		wg.Add(1)
		go exe_cmd(str, wg)
	}
	wg.Wait()
}

func exe_cmd(cmd string, wg *sync.WaitGroup) {
	log.Printf("Command: %s", cmd)
	//parts := strings.Fields(cmd)
	//out, err := exec.Command(parts[0], parts[1]).Output()
	out, err := exec.Command("/bin/sh", "-c", cmd).CombinedOutput()
	if err != nil {
		log.Printf("Error occured: %s", err)
	}
	log.Printf("Output: %s", out)
	wg.Done()
}
