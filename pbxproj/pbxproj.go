package pbxproj

import (
	"sort"

	"github.com/bitly/go-simplejson"
)

// Pbxproj represent project.pbxproj
type Pbxproj struct {
	path               string
	json               *simplejson.Json
	sections           []string
	fileReferences     []FileReference
	nativeTargets      []NativeTarget
	buildFiles         []BuildFile
	sourcesBuildPhases []SourcesBuildPhase
}

// NewPbxproj constructor
func NewPbxproj(path string) *Pbxproj {
	js := convertJSON(path)
	m := js.Get("objects").MustMap()

	return &Pbxproj{
		path,
		js,
		parseSectionNames(m),
		parseFileReferences(m),
		parseNativeTargets(m),
		parseBuildFiles(m),
		parseSourcesBuildPhases(m),
	}
}

// Exists specified section
func (p Pbxproj) Exists(section string) bool {
	return contains(p.SectionNames(), section)
}

// SectionNames return all distinct sorted section names
func (p Pbxproj) SectionNames() []string {
	return p.sections
}

// FileReferencePathNames return file reference path names
func (p Pbxproj) FileReferencePathNames() []string {
	s := []string{}
	for _, f := range p.fileReferences {
		s = append(s, f.path)
	}
	return s
}

// NativeTargetNames return all target names
func (p Pbxproj) NativeTargetNames() []string {
	s := []string{}
	for _, t := range p.nativeTargets {
		s = append(s, t.name)
	}
	return s
}

// BuildFileNames return all build file names
func (p Pbxproj) BuildFileNames() []string {
	s := []string{}
	for _, b := range p.buildFiles {
		if name, found := findFilePath(p.fileReferences, b.fileRef); found {
			s = append(s, name)
		}
	}
	return s
}

// BuildPhaseSourceFileNames return source file for build each target
func (p Pbxproj) BuildPhaseSourceFileNames() map[string][]string {
	m := map[string][]string{}
	for _, s := range p.sourcesBuildPhases {
		t, found := findTargetName(p.nativeTargets, s.id)
		if !found {
			continue
		}
		m[t] = []string{}
		for _, id := range s.files {
			if ref, found := findFileRef(p.buildFiles, id); found {
				if path, found := findFilePath(p.fileReferences, ref); found {
					m[t] = append(m[t], path)
				}
			}
		}
	}
	return m
}

// sectionNames get all distinct sorted section names
// TOOD: modify string to some interface
func parseSectionNames(m map[string]interface{}) []string {
	ss := []string{}

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
