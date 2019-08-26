package leetcode

import "math"

/*
存下所有的丑数，每次从头遍历找到合适的：
超时，时间复杂度：O(n * n * k)
func nthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}
	primes = append([]int{1}, primes...)
	nPrimes := len(primes)
	uglyNumbers := []int{1}

	for  len(uglyNumbers) < n {
		size := len(uglyNumbers)
		lastUglyNumber := uglyNumbers[size - 1]
		minUglyNumber := lastUglyNumber  * primes[nPrimes - 1]
		for i := 0; i < len(uglyNumbers); i ++ {
			for j := 0; j < len(primes); j ++ {
				uNumber := uglyNumbers[i] * primes[j]
				if  uNumber > lastUglyNumber && uNumber < minUglyNumber{
					minUglyNumber = uNumber
				}
			}
		}
		uglyNumbers = append(uglyNumbers, minUglyNumber)
	}
	return uglyNumbers[n - 1]
}
*/


/*
动态规划，本质上是对上述方法的优化，上述方法找下一个超级丑数时需要遍历之前所有的丑数，这一步能做出优化:
下一个丑数的产生一定是由于之前的某个丑数 * [k个丑数列表中的某个素数产生]，因此可以针对每个素数，维护在丑数数组中的位置。
每次选择uglyNumbers[pos(i)] * primes(i)中最小的那个，（i =0...k）。
*/
func nthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}
	primes = append([]int{1}, primes...)
	nPrimes := len(primes)
	uglyNumbers := make([]int, 0)
	uglyNumbers = append(uglyNumbers, 1)
	//与每个prime关联在uglyNumbers中的位置
	mPrimeUglyNumPos := make(map[int]int)
	for i := 0; i < nPrimes; i++ {
		mPrimeUglyNumPos[primes[i]] = 0
	}

	for i := 0; i < n-1; i++ {
		lastUglyNumber := uglyNumbers[i]
		minUglyNumber := math.MaxInt32
		for j := 0; j < nPrimes; j++ {
			pos := mPrimeUglyNumPos[primes[j]]
			tryUglyNumber := uglyNumbers[pos] * primes[j]
			if tryUglyNumber > lastUglyNumber && tryUglyNumber < minUglyNumber {
				minUglyNumber = tryUglyNumber
			}
		}
		uglyNumbers = append(uglyNumbers, minUglyNumber)
		for j := 0; j < nPrimes; j++ {
			pos := mPrimeUglyNumPos[primes[j]]
			if minUglyNumber == uglyNumbers[pos] * primes[j] {
				mPrimeUglyNumPos[primes[j]] += 1
			}
		}
	}
	return uglyNumbers[n-1]
}
