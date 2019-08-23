package leetcode

type Direction int

const Left Direction =  1

const Right Direction =  2

const None  Direction = 0

/*
遇到.统计连续的一共有多少个.
分别考虑这一片垂直区域两侧倾倒的情况，按以下判断顺序
1.左边是R 右侧是L, 则做一半变为R，有一半变为L
2.左边是R，全部变为R
3.右边是L，全部变为L
4.其余情况不改变
 */
func pushDominoes(dominoes string) string {
	if len(dominoes) <= 1 {
		return dominoes
	}
	dominoesBytes := []byte(dominoes)
	for i := 0; i < len(dominoesBytes); {
		if dominoesBytes[i] == '.' {
			startIndex := i
			j := i
			for ;j < len(dominoesBytes) && dominoesBytes[j] == '.'; j ++ {
			}
			endIndex := j - 1
			fill(dominoesBytes, startIndex, endIndex)
			i =  endIndex + 1
		} else {
			i ++
		}
	}
	return string(dominoesBytes)
}

func fill(dominoes []byte, startIndex, endIndex int) {
	left, right := None, None
	if startIndex - 1 >= 0 {
		left = getDir(dominoes, startIndex - 1)
	}
	if endIndex  + 1 < len(dominoes) {
		right = getDir(dominoes, endIndex + 1)
	}
	if left == Right && right == Left {
		for i, j := startIndex, endIndex; i < j; i, j = i + 1, j - 1 {
			dominoes[i] = 'R'
			dominoes[j] = 'L'
		}
	} else if left == Right {
		for i := startIndex; i <= endIndex; i ++ {
			dominoes[i] = 'R'
		}
	} else if right == Left {
		for i := startIndex; i <= endIndex; i ++ {
			dominoes[i] = 'L'
		}
	}
}

func getDir(dominoes []byte, index int) Direction {
	if dominoes[index] == 'R' {
		return Right
	} else {
		return  Left
	}
}
