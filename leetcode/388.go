package leetcode

import "strings"

func lengthLongestPath(input string) int {
	inputs := strings.Split(input,"\n")
	maxLength := 0

	pathStack := make([]string, 0)
	currentLength := 0
	for i := 0; i < len(inputs); i ++ {
		inputs[i] = replace(inputs[i], len(pathStack))


		level := getLevel(inputs[i])
		//去掉前置\t
		s := inputs[i][level:]
		for j := len(pathStack) - 1; j >= level; j -- {
			currentLength -= len(pathStack[j])
		}
		pathStack = pathStack[0 :level]
		//处理文件
		if isFile(inputs[i]) {
			pathLen := currentLength + len(pathStack) + len(inputs[i]) - len(pathStack)
			if pathLen > maxLength {
				maxLength = pathLen
			}
			continue
		}
		//非文件
		if level == len(pathStack) {
			pathStack = append(pathStack, s)
			currentLength += len(s)
		} else {
			pathStack = append(pathStack, s)
			currentLength += len(s)
		}

	}
	return maxLength
}

func getLevel(s string) int {
	if len(s) < 2 {
		return 0
	}
	level := 0

	for i := 0; i < len(s) && s[i] == '\t'; i ++ {
		level ++
	}
	return level
}

func replace(s string, level int) string {
	i := 0
	for ; i + 4 < len(s) && s[i] == ' ' && s[i + 1] == ' ' && s[i + 2] == ' ' && s[i + 3] == ' '; i += 4 {
	}
	cnt := min(i / 4, level)
	result := []byte{}
	for i := 0; i < cnt ;i ++ {
		result = append(result, '\t')
	}
	result = append(result, []byte(s[cnt * 4:])...)
	return string(result)
}

func isFile(s string) bool {
	return strings.Contains(s, ".")
}