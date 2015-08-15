package command

import "github.com/bitly/go-simplejson"

// SourcesBuildPhase represent isa PBXSourcesBuildPhase
type SourcesBuildPhase struct {
	id                                 string
	buildActionMask                    string
	files                              []string
	runOnlyForDeploymentPostprocessing string
}

// parse PBXSourcesBuildPhase
func sourcesBuildPhases(js *simplejson.Json) []SourcesBuildPhase {
	ss := []SourcesBuildPhase{}
	m := js.Get("objects").MustMap()
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
