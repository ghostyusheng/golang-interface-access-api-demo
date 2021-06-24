package identify

import (
	"os/exec"
	"strings"
)

const TEXT = "0"
const GZIP = "1"
const ZIP = "2"

func SimpleIdentify(_m map[string]string) map[string]string {
	cmd := exec.Command("sh", "-c", "file -b "+_m["name"])
	filetype, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	var _type string = strings.Trim(string(filetype), "\n")
	_m["type"] = _type

	if strings.Contains(_type, "text") {
		_m["_type"] = TEXT
	} else if strings.Contains(_type, "gzip") {
		_m["_type"] = GZIP
	} else if strings.Contains(_type, "Zip") {
		_m["_type"] = ZIP
	}

	return _m
}
