package common

import (
	"fmt"
	"os"
)

func GetSizeByFileInfo(fstat os.FileInfo) string {
	var size string
	var _size int64 = fstat.Size()
	if _size >= 0 && _size <= 1024*1024 {
		size = fmt.Sprintf("%.1f", float64(_size)/1024) + "K"
	} else if _size >= 1024*1024 || _size <= 1024*1024*1024 {
		size = fmt.Sprintf("%.1f", float64(_size)/1024/1024) + "M"
	} else {
		size = fmt.Sprintf("%.1f", float64(_size)/1024/1024/1024) + "G"
	}
	return size
}
