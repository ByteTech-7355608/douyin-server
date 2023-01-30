package util

import "encoding/json"

// JSONStr marshals the passed obj interface.
func JSONStr(obj interface{}) string {
	if obj == nil {
		return ""
	}
	bs, _ := json.Marshal(obj)
	return string(bs)
}
