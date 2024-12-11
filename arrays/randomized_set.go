package arrays

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

const (
	MaxUint        = ^uint32(0)
	MinUint32      = 0
	MaxInt32       = int32(MaxUint >> 1)
	MinInt32       = -MaxInt32 - 1
	ValuesRange    = int(MaxInt32) - int(MinInt32) + 1
	CountOfBuckets = 10000
	BucketStep     = int(ValuesRange / CountOfBuckets)
)

type RandomizedSet struct {
	countOfBuckets int
	elements       map[int][]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		countOfBuckets: CountOfBuckets,
		elements:       make(map[int][]int),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	bucket, hasBucket := rs.elements[rs.Hash(val)]
	if hasBucket {
		if rs.isBucketContainsValue(bucket, val) {
			return false
		}
		rs.elements[rs.Hash(val)] = append(bucket, val)
		return true
	} else {
		rs.elements[rs.Hash(val)] = []int{val}
		return true
	}
}

func (rs *RandomizedSet) Remove(val int) bool {
	bucket, hasBucket := rs.elements[rs.Hash(val)]
	if hasBucket {
		for i, v := range bucket {
			if v == val {
				bucket[i] = bucket[len(bucket)-1]
				rs.elements[rs.Hash(val)] = bucket[:len(bucket)-1]
				return true
			}
		}
	}

	return false
}

func (rs *RandomizedSet) GetRandom() int {
	rand.Seed(uint64(time.Now().UnixNano()))
	maxIndex := len(rs.elements)
	fmt.Println(maxIndex)
	index := rand.Intn(maxIndex-0) + maxIndex
	return index
}

func (rs *RandomizedSet) Contains(val int) bool {
	bucket, ok := rs.elements[rs.Hash(val)]
	if ok {
		return rs.isBucketContainsValue(bucket, val)
	}
	return false
}

func (rs *RandomizedSet) isSetContainsBucket(val int) bool {
	_, ok := rs.elements[rs.Hash(val)]
	return ok
}

func (rs *RandomizedSet) isBucketContainsValue(bucket []int, val int) bool {
	for _, v := range bucket {
		if v == val {
			return true
		}
	}
	return false
}

func (rs *RandomizedSet) Hash(val int) int {
	return val / BucketStep
}

func (rs *RandomizedSet) GetValue(val int) int {
	bucket, hasBucket := rs.elements[rs.Hash(val)]
	if hasBucket {
		for _, v := range bucket {
			if v == val {
				return val
			}
		}
	}

	return 505
}

func (rs *RandomizedSet) PrintVals(vals []int) {
	fmt.Printf("\n\n==== RandomizedSet =====")
	for _, v := range vals {
		fmt.Printf("\nval %d, hash %d", rs.GetValue(v), rs.Hash(v))
	}
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
