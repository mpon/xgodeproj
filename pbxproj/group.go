package pbxproj

// Group represent isa PBXGroup
type Group struct {
	id         string
	name       string
	path       string
	children   []string
	sourceTree string
}

func (g Group) isRoot() bool {
	return len(g.name) == 0 && len(g.path) == 0
}

func (g Group) expression() string {
	switch {
	case len(g.name) > 0:
		return g.name
	case len(g.path) > 0:
		return g.path
	default:
		return ""
	}
}

// parse PBXGroup
func parseGroups(m map[string]interface{}) []Group {
	gs := []Group{}

	for id, mm := range m {
		obj := mm.(map[string]interface{})
		for k, v := range obj {
			if k == "isa" && v.(string) == "PBXGroup" {
				g := Group{
					id,
					lookupStr(obj, "name"),
					lookupStr(obj, "path"),
					lookupStrSlices(obj, "children"),
					lookupStr(obj, "sourceTree"),
				}
				gs = append(gs, g)
			}
		}
	}
	return gs
}
