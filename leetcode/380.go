package leetcode

import (
	"math/rand"
)

type RandomizedSet struct {
	dataIdxMap map[int]int
	data []int
}

/** Initialize your data structure here. */
func Constructor2() RandomizedSet {
	return RandomizedSet{
		dataIdxMap: make(map[int]int),
		data:make([]int, 0),
	}
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.dataIdxMap[val]; !ok {
		this.data = append(this.data, val)
		this.dataIdxMap[val] = len(this.data) - 1
		return true
	}
	return  false
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.dataIdxMap[val]; ok {
		n := len(this.data)
		delete(this.dataIdxMap, val)
		if idx != n - 1 {
			this.data[idx] = this.data[n - 1]
			this.dataIdxMap[this.data[idx]] = idx
		}
		this.data = this.data[0: n - 1]
		return true
	}
	return false
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	n := len(this.data)
	randIdx := rand.Intn(n)
	return this.data[randIdx]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
