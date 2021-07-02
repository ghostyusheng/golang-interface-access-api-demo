package identify

import (
	"os/exec"
	"strings"
)

const TEXT = "0"
const GZIP = "1"
const ZIP = "2"

func SimpleIdentify() {
	cmd := exec.Command("sh", "-c", "file -b "+f.Name)
	filetype, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	var _type string = strings.Trim(string(filetype), "\n")
	f.Type = _type

	if strings.Contains(_type, "text") {
		f.Genre = TEXT
	} else if strings.Contains(_type, "gzip") {
		f.Genre = GZIP
	} else if strings.Contains(_type, "Zip") {
		f.Genre = ZIP
	}

}
