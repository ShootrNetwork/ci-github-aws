package main

import "log"

var components = []string{
	"api",
	"services",
	"backoffice",
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Starting...")

	params := parseParams()
	log.Printf("Params: %+v\n", params)

	switch params.Command {

	case cmdTestAndBuild:
		testAndBuild(params)

	case cmdUploadToS3:
		uploadArtifactsToS3(params)

	case cmdDockerBuild:
		dockerBuildComponents(params)

	case cmdDockerTag:
		dockerTagComponents(params)

	case cmdDeploy:
		deployComponents(params)

	default:
		panic("unrecognized command")
	}
}
