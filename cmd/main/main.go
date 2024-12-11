package main

import (
	"cmd/main/arrays"
	"fmt"
)

func main() {
	fmt.Printf("max int32: %d, min int32: %d", arrays.MaxInt32, arrays.MinInt32)
	fmt.Printf("\ndifference %d", arrays.ValuesRange)
	fmt.Printf("\ncount of buckets : %d, bucket size %d", arrays.CountOfBuckets, arrays.BucketStep)
	obj := arrays.Constructor()
	for _, i := range []int{0, 1, -1, 500000, -500000, int(arrays.MaxInt32), int(arrays.MinInt32)} {
		fmt.Printf("\nvalue: %d, bucket: %d", i, obj.Hash(i))
	}
	param_1 := obj.Insert(1)
	obj.Insert(3)
	obj.Remove(3)
	obj.Insert(500000)
	obj.Insert(-500000)
	obj.Insert(int(arrays.MaxInt32))
	obj.Insert(int(arrays.MinInt32))
	obj.Remove(-500000)
	fmt.Printf("\nInsert with ok? %v", param_1)
	obj.PrintVals([]int{1, 3, 500000, -500000, int(arrays.MaxInt32), int(arrays.MinInt32)})
	fmt.Printf("\n=>Random val from set: %d", obj.GetRandom())
	// param_3 := obj.Insert(3)
	// fmt.Println(param_3)
	// param_2 := obj.Remove(2)
	// fmt.Println(param_2)
	// param_4 := obj.GetRandom()
	// fmt.Println(param_4)
}
