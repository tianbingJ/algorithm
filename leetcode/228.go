package leetcode

import "strconv"

type interval struct {
	start int
	end int
}

func summaryRanges(nums []int) []string {

	intervals := []interval{}

	for i := 0; i < len(nums); i ++ {
		count := len(intervals)
		if count == 0 || intervals[count - 1].end  + 1 != nums[i] {
			intervals = append(intervals, interval{nums[i], nums[i]})
		} else  {
			intervals[count - 1].end ++
		}
	}

	result := []string{}
	for i := 0; i < len(intervals); i ++ {
		s := ""
		if intervals[i].start == intervals[i].end {
			s = strconv.Itoa(intervals[i].start)
		} else {
			s = strconv.Itoa(intervals[i].start) + "->" + strconv.Itoa(intervals[i].end)
		}
		result = append(result, s)
	}
	return result
}
