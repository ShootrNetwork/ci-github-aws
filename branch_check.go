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
