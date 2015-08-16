package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/mpon/xgodeproj/pbxproj"
)

// CmdShow for print sections
func CmdShow(c *cli.Context) {

	proj, found := findProjectPath()
	if !found {
		fmt.Println("Not found project.pbxproj file.")
		return
	}

	// get flags
	section := c.String("section")

	// parse pbxproj
	pbxproj := pbxproj.NewPbxproj(proj)

	switch {
	case section == "":
		// show all section names
		for _, s := range pbxproj.SectionNames() {
			fmt.Println(s)
		}
	case !pbxproj.Exists(section):
		fmt.Println(section + " does not exist. try `xgodeproj show` to find section name")
	case section == "PBXFileReference":
		// show file reference paths
		for _, s := range pbxproj.FileReferencePathNames() {
			fmt.Println(s)
		}
	case section == "PBXNativeTarget":
		// show target names
		for _, s := range pbxproj.NativeTargetNames() {
			fmt.Println(s)
		}
	case section == "PBXBuildFile":
		// show build files
		for _, s := range pbxproj.BuildFileNames() {
			fmt.Println(s)
		}
	case section == "PBXSourcesBuildPhase":
		// show build phase source files
		for k, v := range pbxproj.BuildPhaseSourceFileNames() {
			fmt.Println(k)
			fmt.Println("---------")
			for _, s := range v {
				fmt.Println(s)
			}
		}
	default:
		fmt.Println("sorry, not implement parser for the " + section)
	}

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
