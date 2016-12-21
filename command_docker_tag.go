package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func dockerTagComponents(params Params) {
	branchCheck := BranchCheck{params.Config.CurrentConfig}

	if branchCheck.shouldExecuteDockerTag() {
		start := time.Now()
		log.Println("Docker tag start...")

		commit := params.Git.Commit
		tagValue := branchCheck.getDockerTagValue()

		var wg sync.WaitGroup

		for _, component := range components {
			wg.Add(1)
			go pullTagAndPushAsync(component, commit, tagValue, &wg)
		}

		wg.Wait()

		log.Printf("Docker tag done in %s", time.Since(start))

	} else {
		log.Println("Skipping Docker tag")
	}
}

func pullTagAndPushAsync(component string, commit string, tagValue string, wg *sync.WaitGroup) {
	pullTagAndPush(component, commit, tagValue)
	wg.Done()
}

func pullTagAndPush(component string, commit string, tagValue string) {
	commitImage := fmt.Sprintf("fav24/shootr-%s:%s", component, commit)
	tagImage := fmt.Sprintf("fav24/shootr-%s:%s", component, tagValue)

	dockerPull(commitImage)
	dockerTag(commitImage, tagImage)
	dockerPush(tagImage)
	log.Printf("pull and tag for %s:%s done!", component, tagValue)
}
