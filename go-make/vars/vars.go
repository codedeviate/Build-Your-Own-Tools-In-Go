package vars

import (
	"os"
	"regexp"
	"runtime"
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

// Pre-fill the VarList with environment variables
func Init() {
	VarList["PWD"] = os.Getenv("PWD")
	VarList["CURDIR"] = os.Getenv("PWD")
	VarList["HOME"] = os.Getenv("HOME")
	VarList["USER"] = os.Getenv("USER")
	VarList["USERNAME"] = os.Getenv("USER")
	VarList["SHELL"] = os.Getenv("SHELL")
	VarList["PATH"] = os.Getenv("PATH")
	VarList["MAKE"], _ = os.Executable()
	if runtime.GOOS == "windows" {
		VarList["OS"] = "Windows_NT"
	} else {
		VarList["OS"] = ""
	}
}
