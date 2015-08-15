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

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
