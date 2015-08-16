package command

// BuildFile represent isa PBXBuildFile
type BuildFile struct {
	id      string
	fileRef string
}

// find fileRef
func findFileRef(bs []BuildFile, id string) (string, bool) {
	for _, b := range bs {
		if b.id == id {
			return b.fileRef, true
		}
	}
	return "", false
}
