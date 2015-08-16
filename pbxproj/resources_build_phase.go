package pbxproj

// ResourcesBuildPhase represent isa PBXResourcesBuildPhase
type ResourcesBuildPhase struct {
	id                                 string
	buildActionMask                    string
	files                              []string
	runOnlyForDeploymentPostprocessing string
}

// parse PBXResourcesBuildPhase
func parseResourcesBuildPhases(m map[string]interface{}) []ResourcesBuildPhase {
	ss := []ResourcesBuildPhase{}

	for id, mm := range m {
		obj := mm.(map[string]interface{})
		for k, v := range obj {
			if k == "isa" && v.(string) == "PBXResourcesBuildPhase" {
				bp := ResourcesBuildPhase{
					id,
					lookupStr(obj, "buildActionMask"),
					lookupStrSlices(obj, "files"),
					lookupStr(obj, "runOnlyForDeploymentPostprocessing"),
				}
				ss = append(ss, bp)
			}
		}
	}
	return ss
}
