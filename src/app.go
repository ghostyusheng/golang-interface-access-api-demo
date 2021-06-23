package main

import (
	"service/identify"
	"service/table"
)

func main() {
	var finfo map[string]string = identify.BaseInfo()
	table.TablePrint(finfo)
}
