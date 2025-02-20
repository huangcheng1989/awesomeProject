package leetcode

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	s := "pwwkew"
	i := lengthOfLongestSubstring(s)
	t.Log(i)
}

func lengthOfLongestSubstring(s string) int {
	var l int
	for i := 0; i < len(s); i++ {
		m := make(map[int32]struct{})
		var arr []int32
		for j, v := range s {
			if j < i {
				continue
			}

			if _, ok := m[v]; ok {
				break
			}

			m[v] = struct{}{}
			arr = append(arr, v)
		}
		if len(arr) > l {
			l = len(arr)
		}
	}

	return l
}
