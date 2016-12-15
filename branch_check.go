package main

func (self BranchCheck) shouldExecuteTestAndBuild() bool {
	return self.BranchConfig.TestAndBuild
}

func (self BranchCheck) shouldExecuteUploadToS3() bool {
	return self.BranchConfig.UploadToS3
}

func (self BranchCheck) shouldExecuteDockerBuild() bool {
	return self.BranchConfig.DockerBuild
}

func (self BranchCheck) shouldExecuteDockerTag() bool {
	return self.BranchConfig.DockerTag
}

func (self BranchCheck) getDockerTagValue() string {
	return self.BranchConfig.DockerTagValue
}

func (self BranchCheck) shouldDeploy() bool {
	return self.BranchConfig.Deploy
}

func (self BranchCheck) getBackofficeUrl() string {
	return self.BranchConfig.backofficeUrl
}

func (self BranchCheck) getASG() string {
	return self.BranchConfig.ASG
}
