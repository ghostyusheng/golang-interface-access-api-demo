package identify

import (
	"os"
	"strconv"
	wcommon "widget/common"
	wtext "widget/text"
)

type F Finfo
type Finfo struct {
	Name           string
	Genre          string
	Lines          string
	Mode           string
	Size           string
	Type           string
	UserAccess     string
	LastModifyTime string
}

var f = F{}

func BaseInfo() F {
	buildWithFilename()
	SimpleIdentify()
	buildWithStatInfo()
	if f.Genre == TEXT {
		buildWithLineInfo()
	}

	return f
}

func buildWithLineInfo() {
	f.Lines = strconv.Itoa(wtext.CountTextLines(f.Name))
}

func buildWithStatInfo() {
	var fstat, err = os.Stat(f.Name)

	if err != nil {
		panic(err)
	}

	f.Size = wcommon.GetSizeByFileInfo(fstat)
	if fstat.IsDir() == true {
		f.Type = "directory"
	}
	f.Mode = fstat.Mode().String()[1:]
	f.UserAccess = wtext.GetAnalyzedFileMode(fstat)
	f.LastModifyTime = fstat.ModTime().Format("2006-01-02 15:04:05")
}

func buildWithFilename() {
	f.Name = getFilename()
}

func getFilename() string {
	var _f string
	if len(os.Args) == 2 {
		_f = os.Args[1]
		checkExist(_f)
		return _f
	} else {
		panic("demo: ./finfo filename")
	}
}

func checkExist(filename string) {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			panic("file not exist")
		} else {
			panic(err)
		}
	}
}
