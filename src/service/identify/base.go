package identify

import (
	"os"
	"strconv"
	wcommon "widget/common"
	wtext "widget/text"
)

func BaseInfo() map[string]string {
	var finfo = make(map[string]string)
	finfo = buildWithFilename(finfo)
	finfo = SimpleIdentify(finfo)
	finfo = buildWithStatInfo(finfo)
	if finfo["_type"] == TEXT {
		finfo = buildWithLineInfo(finfo)
	}
	return finfo
}

func buildWithLineInfo(_m map[string]string) map[string]string {
	_m["lines"] = strconv.Itoa(wtext.CountTextLines(_m["name"]))
	return _m
}

func buildWithStatInfo(_m map[string]string) map[string]string {
	var fstat, err = os.Stat(_m["name"])

	if err != nil {
		panic(err)
	}

	_m["size"] = wcommon.GetSizeByFileInfo(fstat)
	if fstat.IsDir() == true {
		_m["type"] = "directory"
	}
	_m["mode"] = fstat.Mode().String()[1:]
	_m["user_access"] = wtext.GetAnalyzedFileMode(fstat)
	_m["last_modify_time"] = fstat.ModTime().Format("2006-01-02 15:04:05")
	return _m
}

func buildWithAccessInfo(_m map[string]string) map[string]string {
	_m["access"] = wtext.GetAccessBySysCall(_m["name"])
	return _m
}

func buildWithFilename(_m map[string]string) map[string]string {
	_m["name"] = getFilename()
	return _m
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
