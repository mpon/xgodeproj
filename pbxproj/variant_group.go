package pbxproj

// VariantGroup represent isa PBXVariantGroup
type VariantGroup struct {
	id         string
	name       string
	children   []string
	sourceTree string
}

// parse PBXVariantGroup
func parseVariantGroups(m map[string]interface{}) []VariantGroup {
	gs := []VariantGroup{}

	for id, mm := range m {
		obj := mm.(map[string]interface{})
		for k, v := range obj {
			if k == "isa" && v.(string) == "PBXVariantGroup" {
				g := VariantGroup{
					id,
					lookupStr(obj, "name"),
					lookupStrSlices(obj, "children"),
					lookupStr(obj, "sourceTree"),
				}
				gs = append(gs, g)
			}
		}
	}
	return gs
}
