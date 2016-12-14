package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func copyFile(src string, dst string) {
	// Read all content of src to data
	data, err := ioutil.ReadFile(src)
	check(err)
	// Write data to dst
	err = ioutil.WriteFile(dst, data, 0644)
	check(err)
}

func doWithTimeout(timeout time.Duration, timeBetweenExecutions time.Duration, action func() error) (err error) {
	quitTimer := time.NewTimer(timeout)
	defer quitTimer.Stop()
	for {
		if actErr := action(); actErr != nil {
			err = actErr
			return
		}
		select {
		case <-quitTimer.C:
			err = fmt.Errorf("timed out after %v", timeout)
			return
		default:
			time.Sleep(timeBetweenExecutions)
		}
	}
}
