package pbxproj

// NativeTarget represent isa PBXNativeTarget
type NativeTarget struct {
	id                     string
	buildConfigurationList string
	productReference       string
	productType            string
	productName            string
	buildPhases            []string
	dependencies           []string
	name                   string
	buildRules             []string
}

// parse PBXNativeTarget
func parseNativeTargets(m map[string]interface{}) []NativeTarget {
	ns := []NativeTarget{}

	for id, mm := range m {
		obj := mm.(map[string]interface{})
		for k, v := range obj {
			if k == "isa" && v.(string) == "PBXNativeTarget" {
				nt := NativeTarget{
					id,
					lookupStr(obj, "buildConfigurationList"),
					lookupStr(obj, "productReference"),
					lookupStr(obj, "productType"),
					lookupStr(obj, "productName"),
					lookupStrSlices(obj, "buildPhases"),
					lookupStrSlices(obj, "dependencies"),
					lookupStr(obj, "name"),
					lookupStrSlices(obj, "buildRules"),
				}
				ns = append(ns, nt)
			}
		}
	}
	return ns
}
