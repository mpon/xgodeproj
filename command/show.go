package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/codegangsta/cli"
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
	pbxproj := NewPbxproj(proj)

	switch {
	case section == "":
		// show all section names
		for _, s := range pbxproj.sectionNames {
			fmt.Println(s)
		}
	case !contains(pbxproj.sectionNames, section):
		fmt.Println(section + " does not exist. try `xgodeproj show` to find section name")
	case section == "PBXFileReference":
		// show file reference paths
		for _, f := range pbxproj.fileReferences {
			fmt.Println(f.path)
		}
	case section == "PBXNativeTarget":
		// show native targets
		for _, t := range pbxproj.nativeTargets {
			fmt.Println(t.name)
		}
	case section == "PBXBuildFile":
		// show build files
		for _, bf := range pbxproj.buildFiles {
			if name, found := findFilePath(pbxproj.fileReferences, bf.fileRef); found {
				fmt.Println(name)
			}
		}
	case section == "PBXSourcesBuildPhase":
		// show sources build phases
		for _, s := range pbxproj.sourcesBuildPhases {
			if t, found := findTargetName(pbxproj.nativeTargets, s.id); found {
				fmt.Println(t)
			}
			for _, id := range s.files {
				if ref, found := findFileRef(pbxproj.buildFiles, id); found {
					if p, found := findFilePath(pbxproj.fileReferences, ref); found {
						fmt.Println(" " + p)
					}
				}
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
