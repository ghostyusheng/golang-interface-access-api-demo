package identify

import (
	"os/exec"
	"strings"
)

func SimpleIdentify(_m map[string]string) map[string]string {
	cmd := exec.Command("sh", "-c", "file -b "+_m["name"])
	filetype, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	_m["type"] = strings.Trim(string(filetype), "\n")
	return _m
}
