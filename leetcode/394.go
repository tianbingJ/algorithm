package leetcode

/**
竟然一次过了。
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2 MB, 在所有 Go 提交中击败了100.00%的用户
*/
func decodeString(s string) string {
	//重复次数
	repeatStack := make([]int, 0)
	strStack := make([][]byte, 0)
	repeatStack = append(repeatStack, 1)
	strStack = append(strStack, []byte{})

	for i := 0; i < len(s); i++ {
		//数字
		if isDigit(s[i]) {
			sum := 0
			for ; isDigit(s[i]); i++ {
				sum = sum*10 + int(s[i]) - '0'
			}
			//i指向'['
			repeatStack = append(repeatStack, sum)
			strStack = append(strStack, []byte{})
		} else if s[i] == ']' {
			top := len(repeatStack) - 1
			cnt := repeatStack[top]
			str := []byte{}
			for k := 0; k < cnt; k++ {
				str = append(str, strStack[top]...)
			}
			repeatStack = repeatStack[0:top]
			strStack = strStack[0:top]
			n := len(repeatStack)
			strStack[n-1] = append(strStack[n-1], str...)
		} else {
			staN := len(strStack)
			strStack[staN-1] = append(strStack[staN-1], s[i])
		}
	}
	return string(strStack[0])
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9';
}
