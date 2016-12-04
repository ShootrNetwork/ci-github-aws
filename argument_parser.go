package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

var (
	validCommands = []string{
		"test_and_build",
		"upload_to_s3",
		"docker_build",
		"docker_tag",
		"deploy"}
)

func parseConfig(fileName string) Config {
	var config Config
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		log.Fatal("Error: missing configuration file.\n\n")
	}

	source, err := ioutil.ReadFile(fileName)
	check(err)

	err = yaml.Unmarshal(source, &config)
	check(err)

	return config
}

func parseParams() Params {
	var command = flag.String("c", "", "command [required] -> "+strings.Join(validCommands[:], ", "))
	var branch = flag.String("git-branch", "", "git branch [required] -> needs to be mapped to a ASG in the ci-aws-config.yml")
	var commit = flag.String("git-commit", "", "git commit [required]")
	var isPullRequest = flag.Bool("pr", false, "pull request [optional, default=false]")

	flag.Parse()

	flag.Usage = func() {
		fmt.Printf("usage: ci-aws -c <command> -git-branch <git-branch> [-pr 'false']\n\n")
		fmt.Println("Valid commands: ", validCommands)
		flag.PrintDefaults()
	}

	if *branch == "" {
		log.Fatal("Error: missing branch. To get help, use [--help | -h] option.\n\n")
	}
	if *command == "" {
		log.Fatal("Error: missing command. To get help, use [--help | -h] option.\n\n")
	} else if valid := stringInSlice(*command, validCommands); valid == false {
		log.Fatal("Error: invalid command. To get help, use [--help | -h] option.\n\n")
	}

	var git = Git{Branch: *branch, Commit: *commit, IsPullRequest: *isPullRequest}
	var config = parseConfig("ci-aws-config.yml")
	return Params{Command: *command, Git: git, Config: config}
}
