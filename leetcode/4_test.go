package leetcode

import (
	"sort"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 4}
	nums2 := []int{2, 3}
	f := findMedianSortedArrays(nums1, nums2)
	t.Log(f)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)

	sort.Ints(nums1)
	l := len(nums1)
	var f float64
	if l%2 == 0 {
		f = float64(nums1[l/2-1]+nums1[l/2]) / 2
	} else {
		f = float64(nums1[l/2])
	}

	return f
}
