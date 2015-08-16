package command

import (
	"io/ioutil"
	"os"
	"os/exec"
	"sort"

	"github.com/bitly/go-simplejson"
)

// Pbxproj represent project.pbxproj
type Pbxproj struct {
	path string
	json *simplejson.Json
}

// NewPbxproj constructor
func NewPbxproj(path string) *Pbxproj {
	js := convertJSON(path)
	return &Pbxproj{path, js}
}

// sectionNames get all distinct sorted section names
func (p Pbxproj) sectionNames() []string {
	ss := []string{}
	m := p.json.Get("objects").MustMap()
	for _, mm := range m {
		for k, v := range mm.(map[string]interface{}) {
			if k == "isa" && !contains(ss, v.(string)) {
				ss = append(ss, v.(string))
			}
		}
	}
	sort.Strings(ss)
	return ss
}

// parse PBXFileReference
func (p Pbxproj) fileReferences() []FileReference {
	fs := []FileReference{}
	m := p.json.Get("objects").MustMap()
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

// parse PBXNativeTarget
func (p Pbxproj) nativeTargets() []NativeTarget {
	ns := []NativeTarget{}
	m := p.json.Get("objects").MustMap()
	for id, mm := range m {
		obj := mm.(map[string]interface{})
		for k, v := range obj {
			if k == "isa" && v.(string) == "PBXNativeTarget" {
				nt := NativeTarget{
					id,
					lookupStr(obj, "buildConfigurationList"),
					lookupStr(obj, "productReference"),
					lookupStr(obj, "productType"),
					lookupStr(obj, "productName"),
					lookupStrSlices(obj, "buildPhases"),
					lookupStrSlices(obj, "dependencies"),
					lookupStr(obj, "name"),
					lookupStrSlices(obj, "buildRules"),
				}
				ns = append(ns, nt)
			}
		}
	}
	return ns
}

// parse PBXBuildFile
func (p Pbxproj) buildFiles() []BuildFile {
	bs := []BuildFile{}
	m := p.json.Get("objects").MustMap()
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

// parse PBXSourcesBuildPhase
func (p Pbxproj) sourcesBuildPhases() []SourcesBuildPhase {
	ss := []SourcesBuildPhase{}
	m := p.json.Get("objects").MustMap()
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

// get json from project.pbxproj
func convertJSON(proj string) *simplejson.Json {
	// plutil -convert json -o tmp.json -r project.pbxproj
	tmp := "tmp.json"
	cmd := exec.Command("plutil", "-convert", "json", "-o", tmp, proj)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	// read File to byte type
	rf, err := ioutil.ReadFile(tmp)
	if err != nil {
		panic(err)
	}

	// convert []byte type to json type
	js, err := simplejson.NewJson(rf)
	if err != nil {
		panic(err)
	}
	// temp file removed
	os.Remove(tmp)
	return js
}

// string slices contains string
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// find map string value or default empty string value
func lookupStr(m map[string]interface{}, k string) string {
	if v, found := m[k]; found {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// find map string slices or default empty string slices
func lookupStrSlices(m map[string]interface{}, k string) []string {
	if v, found := m[k]; found {
		a := []string{}
		if vv, ok := v.([]interface{}); ok {
			for _, s := range vv {
				a = append(a, s.(string))
			}
			return a
		}
	}
	return []string{}
}
