package main

import (
	"cmd/main/arrays"
	"fmt"
)

func main() {
	obj := arrays.Constructor()
	param_1 := obj.Insert(1)
	fmt.Println(param_1)
	param_3 := obj.Insert(3)
	fmt.Println(param_3)
	param_2 := obj.Remove(2)
	fmt.Println(param_2)
	param_4 := obj.GetRandom()
	fmt.Println(param_4)
}
