package main

import "io/ioutil"

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
