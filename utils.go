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

func executeWithTimeout(timeout time.Duration, timeBetweenExecutions time.Duration, action func() error) (err error) {
	expirationDate := time.Now().Add(timeout)
	for {
		if actionErr := action(); actionErr != nil {
			err = actionErr
			return
		}
		if time.Now().After(expirationDate) {
			return fmt.Errorf("timed out after %v", timeout)
		} else {
			time.Sleep(timeBetweenExecutions)
		}
	}
}
