package table

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func TablePrint(m map[string]string) {
	const padding = 3
	var lines = make([]string, 20)
	const table_min_width int = 50
	const width int = table_min_width*2 + 1

	var i int = 0
	for k, v := range m {
		lines[i] = k + "\t" + v + "\t"
		i += 1
	}

	w := tabwriter.NewWriter(os.Stdout, table_min_width, 0, padding, ' ', tabwriter.StripEscape|tabwriter.Debug)
	fmt.Fprintln(w, strings.Repeat("-", width))
	for _, s := range lines {
		if s == "" {
			continue
		}
		fmt.Fprintln(w, s)
		fmt.Fprintln(w, strings.Repeat("-", width))
	}
	w.Flush()

}
