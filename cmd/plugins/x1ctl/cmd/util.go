package cmd

import (
	"encoding/json"
	"fmt"
	"strings"
)

// extractVersion walks a JSON blob looking for version-ish fields.
func extractVersion(data []byte) string {
	var v any
	if err := json.Unmarshal(data, &v); err != nil {
		return ""
	}
	var hits []string
	var walk func(any)
	walk = func(val any) {
		switch t := val.(type) {
		case map[string]any:
			for k, v2 := range t {
				lk := strings.ToLower(k)
				if s, ok := v2.(string); ok && (strings.Contains(lk, "firmware") || strings.Contains(lk, "version") || lk == "ver") {
					hits = append(hits, fmt.Sprintf("%s=%s", k, s))
				}
				walk(v2)
			}
		case []any:
			for _, v2 := range t {
				walk(v2)
			}
		}
	}
	walk(v)
	if len(hits) == 0 {
		return ""
	}
	return strings.Join(hits, "; ")
}
