package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"

	"github.com/bitly/go-simplejson"
	"github.com/codegangsta/cli"
)

// FileReference represent isa PBXFileReference
type FileReference struct {
	name              string
	path              string
	lastKnownFileType string
	includeInIndex    string
	explicitFileType  string
	sourceTree        string
}

// CmdShow for print sections
func CmdShow(c *cli.Context) {
	var err error

	// plutil -convert json -o tmp.json -r project.pbxproj
	json := "tmp.json"
	proj := c.Args().First()
	cmd := exec.Command("plutil", "-convert", "json", "-o", json, proj)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	// read File to byte type
	rf, err := ioutil.ReadFile(json)
	if err != nil {
		panic(err)
	}

	// convert []byte type to json type
	js, err := simplejson.NewJson(rf)

	// temp file removed
	os.Remove(json)

	// get section names
	ss := getSections(js)
	for _, s := range ss {
		fmt.Println(s)
	}

	// get file references
	fs := getFileReferences(js)
	for _, f := range fs {
		fmt.Println(f)
	}

}

func getSections(js *simplejson.Json) []string {
	ss := []string{}
	m := js.Get("objects").MustMap()
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

func getFileReferences(js *simplejson.Json) []FileReference {
	fs := []FileReference{}
	m := js.Get("objects").MustMap()
	for _, mm := range m {
		fileRef := mm.(map[string]interface{})
		for k, v := range fileRef {
			if k == "isa" && v.(string) == "PBXFileReference" {
				name := lookupStr(fileRef, "name")
				path := lookupStr(fileRef, "path")
				lastKnowFileType := lookupStr(fileRef, "lastKnowFileType")
				includeInIndex := lookupStr(fileRef, "includeInIndex")
				explicitFileType := lookupStr(fileRef, "explicitFileType")
				sourceTree := lookupStr(fileRef, "sourceTree")
				f := FileReference{
					name, path, lastKnowFileType, includeInIndex,
					explicitFileType, sourceTree}
				fs = append(fs, f)
			}
		}
	}
	return fs
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func lookupStr(m map[string]interface{}, k string) string {
	if v, found := m[k]; found {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}
