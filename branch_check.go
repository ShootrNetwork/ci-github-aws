package main

func (branchCheck BranchCheck) shouldExecuteTestAndBuild() bool {
	return branchCheck.BranchConfig.TestAndBuild
}

func (branchCheck BranchCheck) shouldExecuteUploadToS3() bool {
	return !branchCheck.BranchConfig.IsPullRequest && branchCheck.BranchConfig.UploadToS3
}

func (branchCheck BranchCheck) shouldExecuteDockerBuild() bool {
	return !branchCheck.BranchConfig.IsPullRequest && branchCheck.BranchConfig.DockerBuild
}

func (branchCheck BranchCheck) shouldExecuteDockerTag() bool {
	return !branchCheck.BranchConfig.IsPullRequest && branchCheck.BranchConfig.DockerTag
}

func (branchCheck BranchCheck) getDockerTagValue() string {
	return branchCheck.BranchConfig.DockerTagValue
}

func (branchCheck BranchCheck) shouldDeploy() bool {
	return !branchCheck.BranchConfig.IsPullRequest && branchCheck.BranchConfig.Deploy
}

func (branchCheck BranchCheck) getBackofficeURL() string {
	return branchCheck.BranchConfig.BackofficeURL
}

func (branchCheck BranchCheck) getASG() string {
	return branchCheck.BranchConfig.ASG
}

func (branchCheck BranchCheck) getDeployType() string {
	return branchCheck.BranchConfig.DeployType
}

func (branchCheck BranchCheck) getTargetGroup() string {
	return branchCheck.BranchConfig.TargetGroup
}
