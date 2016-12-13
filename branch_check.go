package main

func (branchCheck BranchCheck) should_execute_test_and_build() bool {
	branchConfigs := &branchCheck.Params.Config.BranchConfig
	for _, branchConfig := range *branchConfigs {
		if branchCheck.Params.Git.Branch == branchConfig.Branch {
			return branchConfig.TestAndBuild
		}
	}
	return branchCheck.Params.Config.Default_test_and_build
}

func (branchCheck BranchCheck) should_execute_upload_to_s3() bool {
	branchConfigs := &branchCheck.Params.Config.BranchConfig
	for _, branchConfig := range *branchConfigs {
		if branchCheck.Params.Git.Branch == branchConfig.Branch {
			return branchConfig.UploadToS3
		}
	}
	return branchCheck.Params.Config.Default_upload_to_s3
}
