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
			return branchConfig.BuildDocker
		}
	}
	return self.Params.Config.Default_build_docker
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
