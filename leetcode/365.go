package leetcode

/**
执行用时 :876 ms, 在所有 Go 提交中击败了6.45%的用户
内存消耗 :42 MB, 在所有 Go 提交中击败了25.00%的用户
搜索的方法。。。。

数学方法是求x 和 y的最大公约数，看是否能被z整除。


 */
func gcd(x, y int) int {
	if x == 0 {
		return y
	}

	return gcd(y %x, x)
}

func canMeasureWater1(x int, y int, z int) bool {
	if z > x + y {
		return false
	}
	if x > y {
		x ,y  = y, x
	}

	if z == 0 {
		return true
	}

	if x == 0 && y == 0 {
		return z == 0
	}

	if x == 0 {
		return z % y == 0
	}

	if y == 0 {
		return z % x == 0
	}

	return z % gcd(x, y) == 0
}


type XY struct {
	x, y int
}

var dp map[int]map[int]bool

func canMeasureWater(x int, y int, z int) bool {
	if z > x+y {
		return false
	}
	if x > y {
		x, y = y, x
	}
	if z == 0 {
		return true
	}
	dp = make(map[int]map[int]bool)
	if !hasValue(0, 0) {
		dp[0][0] = true
	}

	queue := make([]XY, 0)
	queue = append(queue, XY{0, 0})
	for len(queue) > 0 {
		top := queue[0]
		if top.x + top.y == z {
			return true
		}
		queue = queue[1:]
		//清空X
		if top.x > 0 {
			if !hasValue(0, top.y) {
				dp[0][top.y] = true
				queue = append(queue, XY{0, top.y})
			}
		}
		//清空Y
		if top.y > 0 {
			if !hasValue(top.x,0) {
				dp[top.x][0] = true
				queue = append(queue, XY{top.x, 0})
			}
		}
		//装满x
		if !hasValue(x, top.y) {
			dp[x][top.y] = true
			queue = append(queue, XY{x, top.y})
		}
		//装满y
		if !hasValue(top.x, y) {
			dp[top.x][y] = true
			queue = append(queue, XY{top.x, y})
		}

		//x -> y
		xVol := min(top.x, y - top.y)
		if !hasValue(top.x - xVol, top.y + xVol) {
			dp[top.x - xVol][top.y + xVol] = true
			queue = append(queue, XY{top.x - xVol, top.y + xVol})
		}
		//y倒x
		yVol := min(x - top.x, top.y)
		if !hasValue(top.x + yVol, top.y - yVol) {
			dp[top.x + yVol][top.y - yVol] = true
			queue = append(queue, XY{top.x + yVol, top.y - yVol})
		}
	}
	return false
}

func hasValue(x, y int) bool {
	m, ok := dp[x];
	if !ok {
		dp[x] = make(map[int]bool)
		return false
	}
	return m[y]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}