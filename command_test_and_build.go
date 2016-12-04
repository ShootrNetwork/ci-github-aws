package main

import (
	"log"
	"time"
)

func TestAndBuild(params Params) {
	start := time.Now()
	log.Println("Test and build start...")

	log.Printf("Test and build done in %s", time.Since(start))
}
