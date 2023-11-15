package binary_serach_test

import (
	. "binary_serach"
	"fmt"
	"slices"
	"sort"
	"testing"
)

func TestGenerateSortedRandomSlice(t *testing.T) {
	s := GenerateSortedRandomSlice[float64](10, 3.1, 8)
	if len(s) != 10 {
		t.Fatal()
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			t.Fatal("排序並非大到小排序")
		}
	}
}

func ExampleGenerateSortedRandomSlice() {
	s := GenerateSortedRandomSlice(10, 1, 5)
	fmt.Println(len(s))
	// Output:
	// 10
}

func TestFindDuplicates(t *testing.T) {
	arr := []int{1, 3, 2, 5, 5, 10, 8, 1}
	duplicates := FindDuplicates(arr, func(s []int) {
		sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	})
	if !slices.Equal(arr, []int{1, 1, 2, 3, 5, 5, 8, 10}) {
		t.Fatal()
	}
	if !slices.Equal(duplicates, []int{1, 5}) {
		t.Fatal()
	}
}

func ExampleFindDuplicates() {
	arr := []int{1, 1, 2, 3, 5, 5, 8, 10}
	duplicates := FindDuplicates(arr)
	fmt.Println(duplicates)

	// Output:
	// [1 5]
}

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
