package main

func (self BranchCheck) shouldExecuteTestAndBuild() bool {
	return self.BranchConfig.TestAndBuild
}

func (self BranchCheck) shouldExecuteUploadToS3() bool {
	return !self.BranchConfig.IsPullRequest && self.BranchConfig.UploadToS3
}

func (self BranchCheck) shouldExecuteDockerBuild() bool {
	return !self.BranchConfig.IsPullRequest && self.BranchConfig.DockerBuild
}

func (self BranchCheck) shouldExecuteDockerTag() bool {
	return !self.BranchConfig.IsPullRequest && self.BranchConfig.DockerTag
}

func (self BranchCheck) getDockerTagValue() string {
	return self.BranchConfig.DockerTagValue
}

func (self BranchCheck) shouldDeploy() bool {
	return !self.BranchConfig.IsPullRequest && self.BranchConfig.Deploy
}

func (self BranchCheck) getBackofficeUrl() string {
	return self.BranchConfig.BackofficeUrl
}

func (self BranchCheck) getASG() string {
	return self.BranchConfig.ASG
}

func (self BranchCheck) getDeployType() string {
	return self.BranchConfig.DeployType
}
