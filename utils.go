package main

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
