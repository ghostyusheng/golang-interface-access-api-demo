package main

import (
	"service/business"
	"service/identify"
)

func main() {
	var f = identify.BaseInfo()
	business.TablePrint(f)
}
