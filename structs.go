package main

type Params struct {
	Command string
	Git     Git
	Config  Config
}

type Config struct {
	Region string `yaml:"region"`

	Default_deploy         bool   `yaml:"default_deploy"`
	Default_test_and_build bool   `yaml:"default_test_and_build"`
	Default_build_docker   bool   `yaml:"default_build_docker"`
	Default_docker_tag     string `yaml:"default_docker_tag"`
	Default_upload_to_s3   bool   `yaml:"default_upload_to_s3"`

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

type BranchCheck struct {
	Params Params
}
