package exercise_test

import (
	"fmt"
	"go_exercise/algorithms/search"
)

// 在一個已排序的陣列中，找出某一目標值
func Example_binarySearch() {
	arr := []int{1, 1, 2, 3, 5, 5, 8, 10}
	target := 8
	idx := search.BinarySearch(arr, target)
	fmt.Println(idx)

	// Output:
	// 6
}
