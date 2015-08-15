package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/codegangsta/cli"
)

// CmdShow for print sections
func CmdShow(c *cli.Context) {

	proj, found := findProjectPath()
	if !found {
		fmt.Println("Not found project.pbxproj file.")
		return
	}

	json := convertJSON(proj)
	sections := sections(json)
	section := c.String("section")
	fileRefs := fileReferences(json)
	targets := nativeTargets(json)
	sourceBuildPhases := sourcesBuildPhases(json)
	buildFiles := buildFiles(json)

	switch {
	case section == "":
		// show all section names
		for _, s := range sections {
			fmt.Println(s)
		}
	case !contains(sections, section):
		fmt.Println(section + " does not exist. try `xgodeproj show` to find section name")
	case section == "PBXFileReference":
		// show file reference paths
		for _, f := range fileRefs {
			fmt.Println(f.path)
		}
	case section == "PBXNativeTarget":
		// show native targets
		for _, t := range targets {
			fmt.Println(t.name)
		}
	case section == "PBXBuildFile":
		// show build files
		for _, bf := range buildFiles {
			if name, found := findFilePath(fileRefs, bf.fileRef); found {
				fmt.Println(name)
			}
		}
	case section == "PBXSourcesBuildPhase":
		// show sources build phases
		for _, s := range sourceBuildPhases {
			if t, found := findTargetName(targets, s.id); found {
				fmt.Println(t)
			}
			for _, id := range s.files {
				if ref, found := findFileRef(buildFiles, id); found {
					if p, found := findFilePath(fileRefs, ref); found {
						fmt.Println(" " + p)
					}
				}
			}
		}
	default:
		fmt.Println("sorry, not implement parser for the " + section)
	}

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

// get all distinct sorted section names
func sections(js *simplejson.Json) []string {
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

// find project.pbxproj path
func findProjectPath() (projPath string, found bool) {

	// get current directory
	cur, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// find project.pbxproj
	err = filepath.Walk(cur,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				if strings.HasPrefix(info.Name(), ".") {
					return filepath.SkipDir
				}
				return nil
			}
			if filepath.Base(path) == "project.pbxproj" {
				rel, err := filepath.Rel(cur, path)
				if err != nil {
					panic(err)
				}
				projPath = rel
				found = true
			}
			return nil
		})
	return projPath, found
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
