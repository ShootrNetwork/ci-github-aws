package main

type Params struct {
	Command string
	Pem     string
	Git     Git
	Config  Config
}

type Config struct {
	AWS AWS `yaml:"aws"`

	PathInDocker string `yaml:"path_in_docker"`

	Default_deploy           bool   `yaml:"default_deploy"`
	Default_test_and_build   bool   `yaml:"default_test_and_build"`
	Default_docker_build     bool   `yaml:"default_docker_build"`
	Default_docker_tag       bool   `yaml:"default_docker_tag"`
	Default_docker_tag_value string `yaml:"default_docker_tag_value"`
	Default_upload_to_s3     bool   `yaml:"default_upload_to_s3"`

	BranchConfig []struct {
		Branch         string `yaml:"branch"`
		ASG            string `yaml:"asg"`
		Deploy         bool   `yaml:"deploy"`
		TestAndBuild   bool   `yaml:"test_and_build"`
		DockerBuild    bool   `yaml:"docker_build"`
		DockerTag      bool   `yaml:"docker_tag"`
		DockerTagValue string `yaml:"docker_tag_value"`
		UploadToS3     bool   `yaml:"upload_to_s3"`
		backofficeUrl  string `yaml:"backoffice_url"`
	} `yaml:"branch_config"`
}

type AWS struct {
	Region         string `yaml:"region"`
	ActifactBucket string `yaml:"artifact_bucket"`
}

type Git struct {
	Branch        string
	Commit        string
	IsPullRequest bool
}

type BranchCheck struct {
	Params Params
}
