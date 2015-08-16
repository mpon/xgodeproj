package pbxproj

// SourcesBuildPhase represent isa PBXSourcesBuildPhase
type SourcesBuildPhase struct {
	id                                 string
	buildActionMask                    string
	files                              []string
	runOnlyForDeploymentPostprocessing string
}

// parse PBXSourcesBuildPhase
func parseSourcesBuildPhases(m map[string]interface{}) []SourcesBuildPhase {
	ss := []SourcesBuildPhase{}

	for id, mm := range m {
		obj := mm.(map[string]interface{})
		for k, v := range obj {
			if k == "isa" && v.(string) == "PBXSourcesBuildPhase" {
				bp := SourcesBuildPhase{
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
