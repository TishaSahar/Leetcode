package main

import (
	stringsalgs "cmd/main/strings_algs"
	"fmt"
)

func main() {
	haystack := "asdfa"
	needle := "fa"
	res := stringsalgs.StrStr(haystack, needle)
	fmt.Println(res)
}
