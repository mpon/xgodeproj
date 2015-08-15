package command

import "github.com/bitly/go-simplejson"

// BuildFile represent isa PBXBuildFile
type BuildFile struct {
	id      string
	fileRef string
}

// parse PBXBuildFile
func buildFiles(js *simplejson.Json) []BuildFile {
	bs := []BuildFile{}
	m := js.Get("objects").MustMap()
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

// find fileRef
func findFileRef(bs []BuildFile, id string) (string, bool) {
	for _, b := range bs {
		if b.id == id {
			return b.fileRef, true
		}
	}
	return "", false
}
