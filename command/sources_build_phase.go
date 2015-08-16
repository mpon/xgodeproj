package command

// SourcesBuildPhase represent isa PBXSourcesBuildPhase
type SourcesBuildPhase struct {
	id                                 string
	buildActionMask                    string
	files                              []string
	runOnlyForDeploymentPostprocessing string
}
