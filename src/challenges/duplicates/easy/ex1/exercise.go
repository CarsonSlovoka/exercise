package ex1

import (
	. "go_exercise/pkg/utils"
)

// FindDuplicatesWithOptions 針對排序過的資料去找出有哪些數值是重複的 (此為需求1的解答)
// 如果您的資料沒有排序過，可以利用options來幫忙，例如: func(s) {sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })} 注意，這會直接異動原本的arr資料
func FindDuplicatesWithOptions[T Number](arr []T, options ...func(s []T)) (duplicates []T) {
	for _, opt := range options {
		opt(arr)
	}

	n := len(arr)

	// 以下的解法還是針對arr是有排序過
	for i := 0; i < n-1; i++ {
		if arr[i] == arr[i+1] {
			// 有重複

			// 確保這個數字沒有加數到結果之中
			if len(duplicates) == 0 ||
				duplicates[len(duplicates)-1] != arr[i] { // 因為是排序過的資料，所以加入的時候也是按照順序去加，所以有沒有重複加只要判斷最後一個元素即可
				duplicates = append(duplicates, arr[i])
			}
		}
	}
	return duplicates
}

func FindDuplicates[T Number](arr []T) (duplicates []T) {

	// 使用字典來記錄每個元素的個數
	m := make(map[T]int)
	for _, val := range arr {
		if _, exists := m[val]; exists {
			m[val]++
		} else {
			m[val] = 1
		}
	}

	for val, count := range m {
		if count > 1 {
			duplicates = append(duplicates, val)
		}
	}
	return duplicates
}
