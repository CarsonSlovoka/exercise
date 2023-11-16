package ex1_test

import (
	"fmt"
	"go_exercise/challenges/duplicates/easy/ex1"
	"sort"
)

// 找出切片(可能未排序)有重複的數字
func ExampleFindDuplicates() {
	arr := []int{1, 3, 6, 3, 5, 6, 7}
	fmt.Println(ex1.FindDuplicates(arr))
	// Output:
	// [3 6]
}

// 將arr排序，並且找出裡面有重複的數字
func ExampleFindDuplicatesWithOptions() {
	arr := []int{1, 3, 6, 3, 5, 6, 7}
	duplicates := ex1.FindDuplicatesWithOptions(arr, func(s []int) {
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] }) // 會先把arr排序
	})
	fmt.Println(arr)
	fmt.Println(duplicates)
	// Output:
	// [1 3 3 5 6 6 7]
	// [3 6]
}
