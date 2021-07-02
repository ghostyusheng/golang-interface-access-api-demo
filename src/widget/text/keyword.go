package text

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func CountTextKeywords(filename string) string {
	var m map[string]int
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	m = make(map[string]int)
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()
	fd := bufio.NewReader(file)
	for {
		w, err := fd.ReadString(' ')
		if err != nil {
			break
		}
		m[w] += 1
	}

	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	var desc = ""
	var showCount = 5

	for _, kv := range ss {
		if showCount < 0 {
			break
		}
		desc += fmt.Sprintf("%s(%d) ", strings.Trim(kv.Key, " "), kv.Value)
		showCount -= 1
	}
	return desc
}
