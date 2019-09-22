package leetcode

import "math"

/**
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%

实现有点啰嗦
 */

type Equation struct {
	A int
	B int
	K float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	variables := make([]string, 0)         // 变量列表
	variableIdxMap := make(map[string]int) //记录变量的下标

	for i := 0; i < len(equations); i++ {
		for j := 0; j < 2; j++ {
			if _, ok := variableIdxMap[equations[i][j]]; !ok {
				variables = append(variables, equations[i][j])
				variableIdxMap[equations[i][j]] = len(variables) - 1
			}
		}
	}

	valuesMat := make([][]float64, len(variables))
	for i := 0; i < len(variables); i++ {
		valuesMat[i] = make([]float64, len(variables))
	}

	queue := make([]Equation, 0)
	for i := 0; i < len(equations); i++ {
		queue = append(queue, Equation{
			A: variableIdxMap[equations[i][0]],
			B: variableIdxMap[equations[i][1]],
			K: values[i],
		})
	}

	for len(queue) > 0 {
		topEqu := queue[0]
		queue = queue[1:]
		fromIdx := topEqu.A
		toIdx := topEqu.B
		valuesMat[fromIdx][toIdx] = topEqu.K
		valuesMat[toIdx][fromIdx] = 1 / topEqu.K

		//fromIdx
		for i := 0; i < len(variables); i++ {
			idx := variableIdxMap[variables[i]]
			if idx == fromIdx || idx == toIdx {
				continue
			}
			//A/B = K     A/C = K2  => C/B = K / K2
			if !isZero(valuesMat[fromIdx][idx]) && isZero(valuesMat[toIdx][idx]) {
				valuesMat[idx][toIdx] = topEqu.K / valuesMat[fromIdx][idx];
				valuesMat[toIdx][idx] = 1 / valuesMat[idx][toIdx];
				queue = append(queue, Equation{
					A: idx,
					B: toIdx,
					K: valuesMat[idx][toIdx],
				})
			}
			if isZero(valuesMat[fromIdx][idx]) && !isZero(valuesMat[toIdx][idx]) {
				valuesMat[fromIdx][idx] = topEqu.K * valuesMat[toIdx][idx];
				valuesMat[idx][fromIdx] = 1 / valuesMat[fromIdx][idx];
				queue = append(queue, Equation{
					A: fromIdx,
					B: idx,
					K: valuesMat[fromIdx][idx],
				})
			}
		}
	}

	result := make([]float64, len(queries))
	for i := 0; i < len(queries); i++ {
		from, ok1 := variableIdxMap[queries[i][0]]
		to, ok2 := variableIdxMap[queries[i][1]]
		if !ok1 || !ok2 {
			result[i] = -1
			continue
		}
		if from == to {
			result[i] = 1.0
		} else {
			result[i] = valuesMat[from][to];
			if isZero(result[i]) {
				result[i] = -1;
			}
		}
	}
	return result
}

func isZero(d float64) bool {
	return math.Abs(d) < 0.0000001;
}
