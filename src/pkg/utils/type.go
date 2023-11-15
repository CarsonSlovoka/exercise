package utils

type Number interface {
	// T 只能是T格式
	// ~T只要最後形式符合T就算是
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
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
