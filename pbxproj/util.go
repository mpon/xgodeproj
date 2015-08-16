package pbxproj

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/bitly/go-simplejson"
)

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
