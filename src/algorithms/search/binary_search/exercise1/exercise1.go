package binary_serach

import (
	"math/rand"
	"sort"
	"time"
)

type Number interface {
	// T 只能是T格式
	// ~T只要最後形式符合T就算是
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

// GenerateSortedRandomSlice 產生一個slice，他有n個元素，每個元素介於[min, max)之間
func GenerateSortedRandomSlice[T Number](n int, minVal, maxVal T) []T {
	r := rand.New(
		rand.NewSource(time.Now().UnixNano()), // 用時間的數字來當亂樹種子
	)
	s := make([]T, n)

	for i := range s {
		s[i] = minVal + T(r.Int63n(int64(maxVal)))
	}

	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

	return s
}

/* 這其實可行，但是接口還是需要定義Less的方法
// Comparable 是一個界面，用於比較泛型數據
type Comparable[T any] interface {
	Less(b T) bool
}

func SortSlice[T Comparable[T]](s []T) {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[j].Less(s[i]) {
				s[i], s[j] = s[j], s[i] // i, j互調
			}
		}
	}
}
*/

// FindDuplicates 針對排序過的資料去找出有哪些數值是重複的 (此為需求1的解答)
// 此作法類似線性搜尋
// 如果您的資料沒有排序過，可以利用options來幫忙，例如: sort.Slice(s, func(i, j int) bool { return s[i] < s[j] }) 注意，這會直接異動原本的arr資料
func FindDuplicates[T Number](arr []T, options ...func(s []T)) (duplicates []T) {
	for _, opt := range options {
		opt(arr)
	}

	n := len(arr)

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

// BinarySearch 對已排序的s(由小到大)，找尋k是否位於此切片之中，返回其下標值
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
