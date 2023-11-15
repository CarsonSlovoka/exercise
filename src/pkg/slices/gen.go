package slices

import (
	. "go_exercise/pkg/utils"
	"math/rand"
	"sort"
	"time"
)

// GenRandomSlice GenerateSortedRandomSlice 產生一個slice，他有n個元素，每個元素介於[min, max)之間
func GenRandomSlice[T Number](n int, minVal, maxVal T, sorted bool) []T {
	r := rand.New(
		rand.NewSource(time.Now().UnixNano()), // 用時間的數字來當亂樹種子
	)
	s := make([]T, n)

	for i := range s {
		s[i] = minVal + T(r.Int63n(int64(maxVal)))
	}

	if sorted {
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] }) // 小到大排序
	}

	return s
}
