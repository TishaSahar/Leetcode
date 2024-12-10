package arrays

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

type RandomizedSet struct {
	elements map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		elements: make(map[int]int),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if rs.contains(val) {
		return false
	}

	rs.elements[val] = val
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	if rs.contains(val) {
		return false
	}

	delete(rs.elements, val)
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	rand.Seed(uint64(time.Now().UnixNano()))
	maxIndex := len(rs.elements)
	fmt.Println(maxIndex)
	index := rand.Intn(maxIndex-0) + maxIndex
	return rs.pick(index)
}

func (rs *RandomizedSet) contains(val int) bool {
	_, ok := rs.elements[val]
	if ok {
		return true
	}
	return false
}

func (rs *RandomizedSet) pick(index int) int {
	for _, v := range rs.elements {
		if index == 0 {
			return v
		}
		index--
	}
	return -1
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
