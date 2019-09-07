package leetcode

import (
	"testing"
)

func TestLengthLongestPath(t *testing.T) {

	s := "dir\n    file.txt"
	result := lengthLongestPath(s)
	if result != 12 {
		t.Errorf("wrong result %d expected:%d\n", result, 12)
	}
	s = "a\n\tb\n\t\tc.txt\n\taaaa.txt"
	result = lengthLongestPath(s)
	if result != 10 {
		t.Errorf("wrong result %d\n", result)
	}

	s = "dir\n        file.txt"
	result = lengthLongestPath(s)
	if result != 16 {
		t.Errorf("wrong result %d\n", result)
	}
	s = "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
	result = lengthLongestPath(s)
	if result != 32 {
		t.Errorf("wrong result %d\n", result)
	}
}