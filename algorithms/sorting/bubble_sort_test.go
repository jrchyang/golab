package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	input := []int{5, 3, 8, 6, 2}
	expected := []int{2, 3, 5, 6, 8}

	result := BubbleSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BubbleSort(%v) = %v; want %v", input, result, expected)
	}
}

// 基准测试
func BenchmarkBubbleSort(b *testing.B) {
	input := []int{5, 3, 8, 6, 2, 1, 4, 7, 9, 0}
	for i := 0; i < b.N; i++ {
		BubbleSort(input)
	}
}
