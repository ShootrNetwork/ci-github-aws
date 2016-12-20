package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

const (
	cmdTestAndBuild string = "test_and_build"
	cmdUploadToS3   string = "upload_to_s3"
	cmdDockerBuild  string = "docker_build"
	cmdDockerTag    string = "docker_tag"
	cmdDeploy       string = "deploy"
	cmdRunAll       string = "run_all"
)

var (
	validCommands = []string{
		cmdTestAndBuild,
		cmdUploadToS3,
		cmdDockerBuild,
		cmdDockerTag,
		cmdDeploy,
		cmdRunAll}
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
	command := flag.String("c", "", "command [required] -> "+strings.Join(validCommands[:], ", "))
	branch := flag.String("git-branch", "", "git branch [required] -> needs to be mapped to a ASG in the ci-aws-config.yml")
	commit := flag.String("git-commit", "", "git commit [required]")
	isPullRequest := flag.Bool("pr", false, "pull request [optional, default=false]")
	pem := flag.String("pem", "", "pem file [optional]")

	flag.Parse()

	flag.Usage = func() {
		log.Printf("usage: ci-aws -c <command> -git-branch <git-branch> [-pr 'false']\n\n")
		log.Println("Valid commands: ", validCommands)
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

	git := Git{Branch: *branch, Commit: *commit, IsPullRequest: *isPullRequest}
	config := parseConfig("ci-aws-config.yml")

	setCurrentConfig(&config, &git, *branch)

	return Params{Command: *command, Pem: *pem, Git: git, Config: config}
}

func setCurrentConfig(config *Config, git *Git, branch string) {
	var current BranchConfig
	found := false
	for _, branchConfig := range config.AllConfigs {
		if regexp.MustCompile(branchConfig.Branch).MatchString(branch) {
			log.Printf("Found config matching for branch '%s': %s", branch, branchConfig.Branch)
			current = branchConfig
			found = true
			break
		}
	}
	if !found {
		log.Printf("Config for branch '%s' not found, using default settings!", branch)
		current = BranchConfig{
			Deploy:         config.Default_deploy,
			TestAndBuild:   config.Default_test_and_build,
			DockerBuild:    config.Default_docker_build,
			DockerTag:      config.Default_docker_tag,
			DockerTagValue: config.Default_docker_tag_value,
			UploadToS3:     config.Default_upload_to_s3,
		}
	}

	config.CurrentConfig = current
	config.AllConfigs = nil
	config.CurrentConfig.IsPullRequest = git.IsPullRequest
}
