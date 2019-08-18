package leetcode

import "strconv"


func diffWaysToCompute(input string) []int {
	result := []int{}
	for i := 0; i < len(input); i ++ {
		if  input[i] == '+' || input[i] == '-' || input[i] == '*' {
			leftRets := diffWaysToCompute(input[0:i])
			rightRets := diffWaysToCompute(input[i+1:])

			for left := range leftRets {
				for right := range rightRets {
					if input[i] == '+' {
						result = append(result, left + right)
					} else if input[i] == '*' {
						result = append(result, left * right)
					} else if input[i] == '-' {
						result = append(result, left - right)
					}
				}
			}
		}
	}
	if len(result) == 0 {
		v, _ := strconv.Atoi(input)
		result = append(result, v)
	}
	return result
}

