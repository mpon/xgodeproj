package command

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

// find file path
func findFilePath(fs []FileReference, id string) (string, bool) {
	for _, f := range fs {
		if f.id == id {
			return f.path, true
		}
	}
	return "", false
}
