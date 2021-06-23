package identify

import (
	"os"
	"strconv"
)

func BaseInfo() map[string]string {
	var finfo = make(map[string]string)
	finfo = buildWithFilename(finfo)
	finfo = buildWithStatInfo(finfo)
	return finfo
}

func buildWithStatInfo(_m map[string]string) map[string]string {
	var fstat, err = os.Stat(_m["name"])

	if err != nil {
		panic(err)
	}

	_m["size"] = strconv.FormatInt(fstat.Size(), 10)
	if fstat.IsDir() == true {
		_m["type"] = "directory"
	}
	_m["mode"] = fstat.Mode().String()
	_m["modTime"] = fstat.ModTime().Format("2006-01-02 15:04:05")
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
