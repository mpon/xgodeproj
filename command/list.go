package command

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/bitly/go-simplejson"
	"github.com/codegangsta/cli"
)

// CmdList prints list
func CmdList(c *cli.Context) {

	var err error

	// plutil -convert json -o tmp.json -r project.pbxproj
	json := "tmp.json"
	cmd := exec.Command("plutil", "-convert", "json", "-o", json, c.Args()[0])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	// read File to byte type
	rf, err := ioutil.ReadFile(json)
	if err != nil {
		log.Fatal(err)
	}

	// convert []byte type to json type
	js, err := simplejson.NewJson(rf)

	// get classes
	printClasses(js)
	printString(js, "objectVersion")
	printString(js, "archiveVersion")
	printString(js, "rootObject")
	printObjects(js)

	// temp file removed
	os.Remove(json)

}

func printClasses(js *simplejson.Json) {
	fmt.Println("======== classes =========")
	cs := js.Get("classes").MustMap()
	for k, v := range cs {
		fmt.Println(k + " = ")
		fmt.Println(v)
	}
}

func printString(js *simplejson.Json, key string) {
	fmt.Println("======== " + key + " =========")
	v := js.Get(key).MustString()
	fmt.Println(key + " = " + v)
}

func printObjects(js *simplejson.Json) {
	fmt.Println("======== objects =========")
	objects := js.Get("objects").MustMap()
	for _, m := range objects {
		for k, v := range m.(map[string]interface{}) {
			if k == "isa" {
				fmt.Println(k + " = " + v.(string))
			}
		}
	}
}
