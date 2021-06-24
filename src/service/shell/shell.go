package shell

import "os"

func Parse() {
	for k,v := range os.Args {
		println(k, v)
	}
}
