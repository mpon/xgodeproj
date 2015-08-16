package pbxproj

// BuildFile represent isa PBXBuildFile
type BuildFile struct {
	id      string
	fileRef string
}

// parse PBXBuildFile
func parseBuildFiles(m map[string]interface{}) []BuildFile {
	bs := []BuildFile{}

	for id, mm := range m {
		obj := mm.(map[string]interface{})
		for k, v := range obj {
			if k == "isa" && v.(string) == "PBXBuildFile" {
				b := BuildFile{
					id,
					lookupStr(obj, "fileRef"),
				}
				bs = append(bs, b)
			}
		}
	}
	return bs
}
