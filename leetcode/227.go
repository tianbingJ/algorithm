package leetcode

/*
实现一个基本的计算器来计算一个简单的字符串表达式的值。
字符串表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。
*/

func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}
	s = delSpace(s)
	operands := []int{}
	operators := []byte{}
	operand, nextIndex := readOperand(s, 0)
	operands = append(operands, operand)
	for i := nextIndex; i < len(s); {
		switch s[i] {
		case '*', '/':
			operandsCount := len(operands)
			lastValue := operands[operandsCount-1]
			operand, nextIndex = readOperand(s, i+1)
			if s[i] == '*' {
				operands[operandsCount-1] = lastValue * operand
			} else {
				operands[operandsCount-1] = lastValue / operand
			}
			i = nextIndex
		case '+', '-':
			operand, nextIndex = readOperand(s, i+1)
			operators = append(operators, s[i])
			operands = append(operands, operand)
			i = nextIndex
		}
	}
	result := operands[0]
	for i := 0; i < len(operators); i++ {
		if operators[i] == '+' {
			result += operands[i+1]
		} else {
			result -= operands[i+1]
		}
	}
	return result
}

func delSpace(s string) string {
	results := []byte{}
	for i := 0; i < len(s); i ++ {
		if s[i] != ' ' {
			results = append(results, s[i])
		}
	}
	return string(results)
}

func readOperand(s string, start int) (value, nextIndex int) {
	if start >= len(s) {
		return 0, -1
	}
	i := start
	value = 0
	for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
		value = value*10 + int(s[i]-'0')
	}
	return value, i
}
