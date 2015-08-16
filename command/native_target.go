package command

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

func findTargetName(ns []NativeTarget, id string) (string, bool) {
	for _, n := range ns {
		for _, bt := range n.buildPhases {
			if bt == id {
				return n.name, true
			}
		}
	}
	return "", false
}
