package command

import (
	"os"
	"path/filepath"
	"strings"
)

// find project.pbxproj path
func findProjectPath(project string) (projPath string, found bool) {

	isProjectNotSet := len(project) == 0

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
			isPbxproj := filepath.Base(path) == "project.pbxproj"
			isSpecifiedProject := filepath.Base(filepath.Dir(path)) == project+".xcodeproj"
			if isPbxproj && (isProjectNotSet || isSpecifiedProject) {
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
