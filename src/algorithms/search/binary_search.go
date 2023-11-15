package search

import . "go_exercise/pkg/utils"

// BinarySearch 對已排序的s(由小到大)，找尋k是否位於此切片之中，
// 若存在，返回其下標值，
// 若不存在於此切片，則返回-1
func BinarySearch[T Number](s []T, k T, options ...int) int {

	// 空值判斷
	if s == nil || len(s) == 0 {
		return -1
	}

	// 只有一個元素
	if len(s) == 1 {
		if s[0] == k {
			return 0
		}
		return -1
	}

	// 自動決定left, right
	var left, right int
	if len(options) == 0 {
		left = 0
		right = len(s) - 1
	} else {
		left = options[0]
		right = options[1]
	}

	if left > right {
		return -1
	}

	// 二分搜尋

	mid := left + (right-left)/2

	if s[mid] > k {
		return BinarySearch(s, k, 0, mid-1)
	} else if s[mid] < k {
		return BinarySearch(s, k, mid+1, right)
	}

	// 大多數的情況都是不相等，所以放在最後才判斷
	return mid // mid equal k
}
