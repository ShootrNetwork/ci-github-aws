package main

type Params struct {
	Command string
	Git     Git
	Config  Config
}

type Config struct {
	Region       string `yaml:"region"`
	BranchConfig []struct {
		Branch       string `yaml:"branch"`
		ASG          string `yaml:"asg"`
		Deploy       bool   `yaml:"deploy"`
		TestAndBuild bool   `yaml:"test_and_build"`
		BuildDocker  bool   `yaml:"build_docker"`
		DockerTag    string `yaml:"docker_tag"`
		UploadToS3   bool   `yaml:"upload_to_s3"`
	} `yaml:"branch_config"`
}

type Git struct {
	Branch        string
	Commit        string
	IsPullRequest bool
}
