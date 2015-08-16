package pbxproj

// GroupEntry describes xcode project's entry
// it is file or group
type GroupEntry interface {
	IsGroup() bool
	Children(p Pbxproj) []GroupEntry
	Description() string
}

// --- FileReference Implementation ---

// IsGroup returns false
// file is not group
func (f FileReference) IsGroup() bool {
	return false
}

// Children returns empty slice
// file does not have children
func (f FileReference) Children(p Pbxproj) []GroupEntry {
	return []GroupEntry{}
}

// Description returns file path
func (f FileReference) Description() string {
	return f.path
}

// --- Group Implementation ---

// IsGroup returns true
// group is group
func (g Group) IsGroup() bool {
	return true
}

// Children returns GroupEntries
// group have childrens it is group or file
func (g Group) Children(p Pbxproj) []GroupEntry {
	children := []GroupEntry{}
	for _, cID := range g.children {
		if group, found := p.findGroupByID(cID); found {
			children = append(children, group)
		} else if fileRef, found := p.findFileReferenceByID(cID); found {
			children = append(children, fileRef)
		}
	}
	return children
}

// Description returns group expression
func (g Group) Description() string {
	return g.expression()
}
