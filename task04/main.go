package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	items := map[string]map[string][]string{
		"Reader1": {
			"Books":      {"5"},
			"Pereodical": {"2"},
		},
		"Reader2": {
			"Books":      {"2"},
			"Pereodical": {"4"},
		},
		"Reader3": {
			"Books":      {"4"},
			"Pereodical": {"22"},
		},
		"Reader4": {
			"Books":      {"0"},
			"Pereodical": {"22"},
		},
		"Reader5": {
			"Books":      {"22"},
			"Pereodical": {"0"},
		},
		"Reader6": {
			"Books":      {"0"},
			"Pereodical": {"0"},
		},
	}

	keys := make([]string, 0, len(items))
	for k := range items {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var f *int
	var ff int
	f = &ff

	for _, i := range keys {
		if strings.Join(items[i]["Books"], " ") != "0" || strings.Join(items[i]["Pereodical"], " ") != "0" {
			ff++

		}

	}
	fmt.Println(*f)

	v := make([]int, 0, len(items)*2)

	for _, i := range keys {
		for m := range items[i] {
			for p := range items[i][m] {

				c, _ := strconv.Atoi(items[i][m][p])
				v = append(v, c)

			}

		}

	}
	fmt.Print(v)
	vv := make([]int, len(items))
	m := 1
	p := 0

	for k := 0; k < len(items); k++ {

		vv[k] = v[p] + v[m]
		p = p + 2
		m = p + 1

		//fmt.Println(vv[k])
	}
	fmt.Println(vv)
	for i := 0; i < len(items); i++ {

		fmt.Printf("У %s на руках %v книг \n", keys[i], vv[i])

	}

}
