package leetcode

import "math/rand"

type Solution5 struct {
	nums[]int
}


func Constructor5(nums []int) Solution5 {
	return Solution5{nums:nums}
}


/** Resets the array to its original configuration and return it. */
func (this *Solution5) Reset() []int {
	return this.nums
}


/** Returns a random shuffling of the array. */
func (this *Solution5) Shuffle() []int {
	n := len(this.nums)
	data := make([]int, n)
	copy(data, this.nums)
	result := make([]int, n)
	for i := 0; i < n; i ++ {
		randIdx := rand.Intn(n - i)
		result[i] = data[randIdx]
		if randIdx != n - i - 1 {
			data[randIdx] = data[n - i - 1]
		}
	}
	return result
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */