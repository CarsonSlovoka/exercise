package search_test

import (
	"fmt"
	. "go_exercise/algorithms/search"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	if BinarySearch(nil, 200) != -1 {
		t.Fail()
	}

	if BinarySearch([]float32{}, 200) != -1 {
		t.Fail()
	}

	if BinarySearch([]float32{5}, 8.3) != -1 {
		t.Fail()
	}

	if BinarySearch([]float32{3.14}, 3.14) != 0 {
		t.Fail()
	}

	// 在最左邊
	if BinarySearch([]int{1, 2, 3}, 1) != 0 {
		t.Fail()
	}

	// 剛好在中間
	if BinarySearch([]int{1, 2, 3}, 2) != 1 {
		t.Fail()
	}

	// 在最右邊
	if BinarySearch([]int{1, 2, 3}, 3) != 2 {
		t.Fail()
	}

	// 左區域
	if BinarySearch([]float32{5, 8.3, 9, 10, 11, 11, 15, 15}, 8.3) != 1 {
		t.Fail()
	}

	// 右區域
	if BinarySearch([]float32{5, 8.3, 9, 10, 11, 11, 15, 15}, 15) != 6 {
		t.Fail()
	}

	// 只對下標值2~5之間去搜尋: 沒找到
	if idx := BinarySearch([]float32{5, 8.3, 9, 10, 11, 11, 15, 15}, 15, 2, 5); idx != -1 {
		t.Error(idx)
	}

	// 只對下標值2~5之間去搜尋: 找到
	if BinarySearch([]float32{5, 8.3, 9, 10, 11, 11, 15, 15}, 11, 2, 5) != 4 {
		t.Fail()
	}

	// left > right
	if BinarySearch([]float32{5, 8.3, 9, 10, 11, 11, 15, 15}, 15, 0, -5) != -1 {
		t.Error("left > right")
	}
}

func ExampleBinarySearch() {
	arr := []float32{1, 2, 3, 5, 5, 8.3, 10}
	fmt.Println(BinarySearch(arr, 8.3))
	// Output:
	// 5
}
