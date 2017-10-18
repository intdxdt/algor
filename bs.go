package algor

import (
	"github.com/intdxdt/cmp"
)

//index of item in array
func IndexOf(array []interface{}, key interface{}, comparator cmp.Compare, ij ...int) int {
	var idx int
	if len(ij) > 0 {
		idx = BS(array, key, comparator, ij[0], ij[1])
	} else {
		idx = BS(array, key, comparator)
	}
	if idx < 0 {
		idx = -idx - 1
	}
	return idx
}

//binary search assumes the array is sorted
func BS(array []interface{}, key interface{}, comparator cmp.Compare, ij ...int) int {
	low  := 0
	high := len(array) - 1
	if len(ij) > 0 {
		low = ij[0]
		high = ij[1]
	}

	for low <= high {
		mid := low + ((high - low) >> 1)
		v := comparator(array[mid], key)
		if v < 0 {
			low = mid + 1 //too low
		} else if v > 0 {
			high = mid - 1 //too high
		} else {
			return mid //found
		}
	}
	return -(low + 1) //key not found
}
