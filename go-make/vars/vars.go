package vars

import (
	"os"
	"regexp"
)

var VarList map[string]interface{} = make(map[string]interface{})

func EvalVar(input string) string {
	re := regexp.MustCompile(`\$\([a-zA-Z0-9_-]+\)`)
	val := re.ReplaceAllStringFunc(input, func(match string) string {
		key := match[2 : len(match)-1]
		if _, ok := VarList[key]; !ok {
			if len(os.Getenv(key)) > 0 {
				return os.Getenv(key)
			}
			return match
		}
		return VarList[key].(string)
	})
	return val
}
