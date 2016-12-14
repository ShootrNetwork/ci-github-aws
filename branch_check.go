package main

func (self BranchCheck) shouldExecuteTestAndBuild() bool {
	branchConfigs := &self.Params.Config.BranchConfig
	for _, branchConfig := range *branchConfigs {
		if self.Params.Git.Branch == branchConfig.Branch {
			return branchConfig.TestAndBuild
		}
	}
	return self.Params.Config.Default_test_and_build
}

func (self BranchCheck) shouldExecuteUploadToS3() bool {
	branchConfigs := &self.Params.Config.BranchConfig
	for _, branchConfig := range *branchConfigs {
		if self.Params.Git.Branch == branchConfig.Branch {
			return branchConfig.UploadToS3
		}
	}
	return self.Params.Config.Default_upload_to_s3
}

func (self BranchCheck) shouldExecuteDockerBuild() bool {
	branchConfigs := &self.Params.Config.BranchConfig
	for _, branchConfig := range *branchConfigs {
		if self.Params.Git.Branch == branchConfig.Branch {
			return branchConfig.DockerBuild
		}
	}
	return self.Params.Config.Default_docker_build
}

func (self BranchCheck) shouldExecuteDockerTag() bool {
	branchConfigs := &self.Params.Config.BranchConfig
	for _, branchConfig := range *branchConfigs {
		if self.Params.Git.Branch == branchConfig.Branch {
			return branchConfig.DockerTag
		}
	}
	return self.Params.Config.Default_docker_tag
}

func (self BranchCheck) getDockerTagValue() string {
	branchConfigs := &self.Params.Config.BranchConfig
	for _, branchConfig := range *branchConfigs {
		if self.Params.Git.Branch == branchConfig.Branch {
			return branchConfig.DockerTagValue
		}
	}
	return self.Params.Config.Default_docker_tag_value
}

func (self BranchCheck) shouldDeploy() bool {
	branchConfigs := &self.Params.Config.BranchConfig
	for _, branchConfig := range *branchConfigs {
		if self.Params.Git.Branch == branchConfig.Branch {
			return branchConfig.Deploy
		}
	}
	return self.Params.Config.Default_deploy
}

func (self BranchCheck) getBackofficeUrl() string {
	var retVal string
	branchConfigs := &self.Params.Config.BranchConfig
	for _, branchConfig := range *branchConfigs {
		if self.Params.Git.Branch == branchConfig.Branch {
			retVal = branchConfig.backofficeUrl
		}
	}
	return retVal
}

func (self BranchCheck) getASG() string {
	var retVal string
	branchConfigs := &self.Params.Config.BranchConfig
	for _, branchConfig := range *branchConfigs {
		if self.Params.Git.Branch == branchConfig.Branch {
			retVal = branchConfig.ASG
		}
	}
	return retVal
}
