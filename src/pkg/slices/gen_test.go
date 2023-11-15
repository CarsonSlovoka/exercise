package slices_test

import (
	"go_exercise/pkg/slices"
	"testing"
)

func TestGenRandomSlice(t *testing.T) {
	s := slices.GenRandomSlice[float64](10, 3.1, 8, true)
	if len(s) != 10 {
		t.Fail()
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			t.Error("排序並非大到小排序")
		}
	}
}
