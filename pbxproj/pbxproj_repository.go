package pbxproj

import "sort"

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

// find BuildFile
func (p Pbxproj) findBuildFileByID(id string) (BuildFile, bool) {
	for _, b := range p.buildFiles {
		if b.id == id {
			return b, true
		}
	}
	return BuildFile{}, false
}

// find FireReference
func (p Pbxproj) findFileReferenceByID(id string) (FileReference, bool) {
	for _, f := range p.fileReferences {
		if f.id == id {
			return f, true
		}
	}
	return FileReference{}, false
}

// find target
func (p Pbxproj) findNativeTargetByID(id string) (NativeTarget, bool) {
	for _, nt := range p.nativeTargets {
		for _, bid := range nt.buildPhases {
			if bid == id {
				return nt, true
			}
		}
	}
	return NativeTarget{}, false
}

// find variant group
func (p Pbxproj) findVariantGroupByID(id string) (VariantGroup, bool) {
	for _, vg := range p.variantGroups {
		if vg.id == id {
			return vg, true
		}
	}
	return VariantGroup{}, false
}

// find group
func (p Pbxproj) findGroupByID(id string) (Group, bool) {
	for _, group := range p.groups {
		if group.id == id {
			return group, true
		}
	}
	return Group{}, false
}
