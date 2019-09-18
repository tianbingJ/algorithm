package leetcode

func trap1(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	leftH := make([]int, n)
	rightH := make([]int, n)

	leftH[0] = height[0]
	for i := 1; i < n; i ++ {
		leftH[i] = max(leftH[i - 1], height[i])
	}
	rightH[n - 1] = height[n - 1]
	for i := n - 2; i >= 0; i -- {
		rightH[i] = max(rightH[i + 1], height[i])
	}
	res := 0
	for i := 1; i < n - 1; i ++ {
		v := min(leftH[i], rightH[i]) - height[i]
		if v > 0  {
			res += v;
		}
	}
	return res
}

/**
执行用时 :4 ms, 在所有 Go 提交中击败了93.87%的用户
 */
func trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	lMax := height[0]
	rMax := height[n -1]

	lIdx := 1
	rIdx := n - 2
	res := 0
	for lIdx <= rIdx  {
		if lMax > rMax {
			v := min(lMax, rMax) - height[rIdx]
			if v > 0 {
				res += v
			}
			rMax = max(rMax, height[rIdx])
			rIdx --
		} else {
			v := min(lMax, rMax) - height[lIdx]
			if v > 0 {
				res += v
			}
			lMax = max(lMax, height[lIdx])
			lIdx ++
		}
	}
	return res
}