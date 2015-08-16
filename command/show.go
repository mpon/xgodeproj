package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/mpon/xgodeproj/pbxproj"
)

// CmdShow for print sections
func CmdShow(c *cli.Context) {

	// get project flag
	project := c.String("project")

	// find project.pbxproj path
	projPath, found := findProjectPath(project)
	if !found {
		fmt.Println("Not found project.pbxproj file.")
		return
	}

	// get flags
	section := c.String("section")
	isSectionNotSet := len(section) == 0

	// parse pbxproj
	proj := pbxproj.NewPbxproj(projPath)

	switch {
	case isSectionNotSet:
		// show all section names
		for _, s := range proj.SectionNames() {
			fmt.Println(s)
		}
	case !proj.Exists(section):
		fmt.Println(section + " does not exist. try `xgodeproj show` to find section name")
	case section == "PBXFileReference":
		// show file reference paths
		for _, s := range proj.FileReferencePathNames() {
			fmt.Println(s)
		}
	case section == "PBXNativeTarget":
		// show target names
		for _, s := range proj.NativeTargetNames() {
			fmt.Println(s)
		}
	case section == "PBXBuildFile":
		// show build files
		for _, s := range proj.BuildFileNames() {
			fmt.Println(s)
		}
	case section == "PBXSourcesBuildPhase":
		// show build phase source files
		for k, v := range proj.BuildPhaseSourceFileNames() {
			fmt.Println(k)
			fmt.Println("---------")
			for _, s := range v {
				fmt.Println(s)
			}
		}
	case section == "PBXResourcesBuildPhase":
		// show build phase resource files
		for k, v := range proj.BuildPhaseResourceFileNames() {
			fmt.Println(k)
			fmt.Println("---------")
			for _, s := range v {
				fmt.Println(s)
			}
		}
	case section == "PBXVariantGroup":
		// show variant groups
		for _, s := range proj.VariantGroupNames() {
			fmt.Println(s)
		}
	case section == "PBXGroup":
		// show group tree
		proj.Walk(func(entry pbxproj.GroupEntry, level int) {
			if level == 0 {
				return
			}
			for i := 1; i < level; i++ {
				fmt.Print("  ")
			}
			if entry.IsGroup() {
				fmt.Println("+ " + entry.Description())
			} else {
				fmt.Println("  " + entry.Description())
			}
		})

	default:
		fmt.Println("sorry, not implement parser for the " + section)
	}

}
