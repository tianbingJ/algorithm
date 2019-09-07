package leetcode

func validUtf8(data []int) bool {

	for i := 0; i < len(data); {
		//1字节
		if (data[i]&0X80) == 0 {
			i += 1
			continue
		}
		//>1字节
		var n int
		if (data[i] & 0X80) == 0X80 {
			n = 1
		}
		if (data[i] & 0XC0) == 0XC0 {
			n = 2
		}
		if (data[i] & 0XE0) == 0XE0 {
			n = 3
		}
		if (data[i] & 0XF0) == 0XF0 {
			n = 4
		}
		if n == 1 {
			return false
		}
		//验证n + 1位
		if (data[i] & (1 << uint(8-n - 1))) != 0 {
			return false
		}
		for j := i + 1; j < i+n; j++ {
			if j == len(data) {
				return false
			}
			if data[j]&0X80 != 0X80 {
				return false
			}
		}
		i = i + n
	}
	return true
}
