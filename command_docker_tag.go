package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func dockerTagComponents(params Params) {
	branchCheck := BranchCheck{Params: params}

	if branchCheck.shouldExecuteDockerTag() {
		start := time.Now()
		log.Println("Docker tag start...")

		commit := params.Git.Commit
		tagValue := branchCheck.getDockerTagValue()

		wg := new(sync.WaitGroup)

		for _, component := range components {
			go func(component string, commit string, tagValue string) {
				wg.Add(1)
				pullTagAndPush(component, commit, tagValue)
				wg.Done()
			}(component, commit, tagValue)
		}

		wg.Wait()

		log.Printf("Docker tag done in %s", time.Since(start))

	} else {
		log.Println("Skipping Docker tag")
	}
}

func pullTagAndPush(component string, commit string, tagValue string) {
	commitImage := fmt.Sprintf("fav24/shootr-%s:%s", component, commit)
	tagImage := fmt.Sprintf("fav24/shootr-%s:%s", component, tagValue)

	dockerPull(commitImage)
	dockerTag(commitImage, tagImage)
	dockerPush(tagImage)
}
