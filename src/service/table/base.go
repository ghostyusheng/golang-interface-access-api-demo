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

	var i int = 0
	for k, v := range m {
		lines[i] = k + "\t" + v + "\t"
		i += 1
	}

	w := tabwriter.NewWriter(os.Stdout, 25, 0, padding, ' ', tabwriter.StripEscape|tabwriter.Debug)
	fmt.Fprintln(w, strings.Repeat("-", 25*2+1))
	for _, s := range lines {
		if s == "" {
			continue
		}
		fmt.Fprintln(w, s)
		fmt.Fprintln(w, strings.Repeat("-", 25*2+1))
	}
	s := `word frequency           | the:23 of:22 to:15      |`
	fmt.Fprintln(w, s)
	fmt.Fprintln(w, strings.Repeat("-", 25*2+1))
	w.Flush()

}
