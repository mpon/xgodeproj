package command

import "github.com/bitly/go-simplejson"

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
func fileReferences(js *simplejson.Json) []FileReference {
	fs := []FileReference{}
	m := js.Get("objects").MustMap()
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
