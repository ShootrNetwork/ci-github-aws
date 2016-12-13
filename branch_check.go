package main

func (self BranchCheck) should_execute_test_and_build() bool {
	branchConfigs := &self.Params.Config.BranchConfig
	for _, branchConfig := range *branchConfigs {
		if self.Params.Git.Branch == branchConfig.Branch {
			return branchConfig.TestAndBuild
		}
	}
	return self.Params.Config.Default_test_and_build
}

func (self BranchCheck) should_execute_upload_to_s3() bool {
	branchConfigs := &self.Params.Config.BranchConfig
	for _, branchConfig := range *branchConfigs {
		if self.Params.Git.Branch == branchConfig.Branch {
			return branchConfig.UploadToS3
		}
	}
	return self.Params.Config.Default_upload_to_s3
}

func (self BranchCheck) should_execute_docker_build() bool {
	branchConfigs := &self.Params.Config.BranchConfig
	for _, branchConfig := range *branchConfigs {
		if self.Params.Git.Branch == branchConfig.Branch {
			return branchConfig.BuildDocker
		}
	}
	return self.Params.Config.Default_build_docker
}
