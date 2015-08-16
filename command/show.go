package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/mpon/xgodeproj/pbxproj"
)

// CmdShow for print sections
func CmdShow(c *cli.Context) {

	// find project.pbxproj path
	proj, found := findProjectPath()
	if !found {
		fmt.Println("Not found project.pbxproj file.")
		return
	}

	// get flags
	section := c.String("section")
	isSectionNotSet := section == ""

	// parse pbxproj
	pbxproj := pbxproj.NewPbxproj(proj)

	switch {
	case isSectionNotSet:
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
	case section == "PBXResourcesBuildPhase":
		// show build phase resource files
		for k, v := range pbxproj.BuildPhaseResourceFileNames() {
			fmt.Println(k)
			fmt.Println("---------")
			for _, s := range v {
				fmt.Println(s)
			}
		}
	case section == "PBXVariantGroup":
		// show variant groups
		for _, s := range pbxproj.VariantGroupNames() {
			fmt.Println(s)
		}
	default:
		fmt.Println("sorry, not implement parser for the " + section)
	}

}
