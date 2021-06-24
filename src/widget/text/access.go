package text

import (
	"os"
	"syscall"
	"util"
)

func GetAccessBySysCall(name string) string {
	var access string
	err := syscall.Access(name, syscall.O_RDWR)
	if err == nil {
		access = "writable"
	} else {
		access = "non-writable"
	}
	return access
}

func GetAnalyzedFileMode(fstat os.FileInfo) string {
	var mode string = fstat.Mode().String()[1:]
	info, _ := fstat.Sys().(*syscall.Stat_t)
	var fUid int = int(info.Uid)
	var fGid int = int(info.Gid)
	var sUid int = os.Getuid()
	var sGid int = os.Getgid()

	var user_access string = translateModeRWX(mode, sUid, fUid, sGid, fGid)
	return user_access
}

func translateModeRWX(mode string, sUid int, fUid int, sGid int, fGid int) string {
	var m1 string = mode[0:3]
	var m2 string = mode[3:6]
	var m3 string = mode[6:9]
	var flag int = 0
	_m := make(map[string]string, 3)
	if sUid == fUid {
		flag = 1
		for _, i := range m1 {
			_m[string(i)] = "1"
		}
	}
	if sGid == fGid {
		flag = 1
		for _, i := range m2 {
			_m[string(i)] = "1"
		}
	}
	if flag == 0 {
		for _, i := range m3 {
			_m[string(i)] = "1"
		}
	}
	delete(_m, "-")
	var _mstr string = util.CreateKeyValuePairs(_m)
	return _mstr
}
