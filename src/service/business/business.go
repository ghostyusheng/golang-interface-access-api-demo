package business

import (
	"service/identify"
	"service/table"
)

func TablePrint(f identify.F) {
	var m = transferFinfoToOrderedMap(f)
	table.TablePrint(m)
}

func transferFinfoToOrderedMap(f identify.F) map[string]string {
	var m = make(map[string]string)
	m["name"] = f.Name
	m["size"] = f.Size
	m["type"] = f.Type
	m["user access"] = f.UserAccess
	m["last modify time"] = f.LastModifyTime
	m["mode"] = f.Mode
	m["lines"] = f.Lines
	m["keywords"] = f.Keywords
	return m
}
