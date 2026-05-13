package webdocs

import "fmt"

func init() {
	tplFuncs["dict"] = dict
}

// dict packs alternating key/value pairs into a map for template use.
// Pattern: {{template "node" dict "Node" . "MetadataKeys" $keys}}.
func dict(kv ...any) (map[string]any, error) {
	if len(kv)%2 != 0 {
		return nil, fmt.Errorf("dict: odd argument count")
	}
	out := make(map[string]any, len(kv)/2)
	for i := 0; i < len(kv); i += 2 {
		k, ok := kv[i].(string)
		if !ok {
			return nil, fmt.Errorf("dict: key at position %d is not a string", i)
		}
		out[k] = kv[i+1]
	}
	return out, nil
}
