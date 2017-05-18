package main

type Params struct {
	Command string
	Pem     string
	Git     Git
	Config  Config
}

type Config struct {
	AWS AWS `yaml:"aws"`

	PathInDocker         string `yaml:"path_in_docker"`
	DockerBuildImageName string `yaml:"docker_build_image_name"`
	MvnBuildArgs         string `yaml:"mvn_build_args"`

	Default_deploy           bool   `yaml:"default_deploy"`
	Default_test_and_build   bool   `yaml:"default_test_and_build"`
	Default_docker_build     bool   `yaml:"default_docker_build"`
	Default_docker_tag       bool   `yaml:"default_docker_tag"`
	Default_docker_tag_value string `yaml:"default_docker_tag_value"`
	Default_upload_to_s3     bool   `yaml:"default_upload_to_s3"`
	Default_deploy_type      string `yaml:"default_deploy_type"`

	Components    []Component    `yaml:"components"`
	AllConfigs    []BranchConfig `yaml:"branch_config"`
	CurrentConfig BranchConfig
}

type Component struct {
	JarName        string `yaml:"jar_name"`
	JarPath        string `yaml:"jar_path"`
	DockerFilePath string `yaml:"docker_file_path"`
	DockerImage    string `yaml:"docker_image"`
}

type BranchConfig struct {
	Branch         string `yaml:"branch"`
	ASG            string `yaml:"asg"`
	Deploy         bool   `yaml:"deploy"`
	DeployType     string `yaml:"deploy_type"`
	TestAndBuild   bool   `yaml:"test_and_build"`
	DockerBuild    bool   `yaml:"docker_build"`
	DockerTag      bool   `yaml:"docker_tag"`
	DockerTagValue string `yaml:"docker_tag_value"`
	UploadToS3     bool   `yaml:"upload_to_s3"`
	BackofficeURL  string `yaml:"backoffice_url"`
	IsPullRequest  bool
}

type AWS struct {
	Region         string `yaml:"region"`
	ActifactBucket string `yaml:"artifact_bucket"`
	ArtifactFolder string `yaml:"artifact_folder"`
}

type Git struct {
	Branch        string
	Commit        string
	IsPullRequest bool
}

type BranchCheck struct {
	BranchConfig BranchConfig
}
