package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareV(version1, version2 string) int {
	res := 0

	p := 0
	verArr1 := strings.Split(version1, ".")
	verArr2 := strings.Split(version2, ".")

	for p < len(verArr1) || p < len(verArr2) {
		v1, v2 := 0, 0
		if p < len(verArr1) {
			v1, _ = strconv.Atoi(verArr1[p])
		}
		if p < len(verArr2) {
			v2, _ = strconv.Atoi(verArr2[p])
		}

		if v1 > v2 {
			return 1
		}

		if v1 < v2 {
			return -1
		}

		p++
	}

	return res
}

func main() {
	v := []string{"1.0", "1.0.1", "2.1", "0.1"}
	for _, s := range v {
		res := compareV(v[0], s)
		fmt.Println(v[0], " vs ", s, ":", res)
	}

}
