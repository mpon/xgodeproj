package pbxproj

// FileReference represent isa PBXFileReference
type FileReference struct {
	id                string
	name              string
	path              string
	lastKnownFileType string
	includeInIndex    string
	explicitFileType  string
	sourceTree        string
}

// parse PBXFileReference
func parseFileReferences(m map[string]interface{}) []FileReference {
	fs := []FileReference{}

	for id, mm := range m {
		obj := mm.(map[string]interface{})
		for k, v := range obj {
			if k == "isa" && v.(string) == "PBXFileReference" {
				f := FileReference{
					id,
					lookupStr(obj, "name"),
					lookupStr(obj, "path"),
					lookupStr(obj, "lastKnownFileType"),
					lookupStr(obj, "includeInIndex"),
					lookupStr(obj, "explicitFileType"),
					lookupStr(obj, "sourceTree"),
				}
				fs = append(fs, f)
			}
		}
	}
	return fs
}

// find file path
func findFilePath(fs []FileReference, id string) (string, bool) {
	for _, f := range fs {
		if f.id == id {
			return f.path, true
		}
	}
	return "", false
}
