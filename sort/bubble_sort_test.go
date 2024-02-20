package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{2, 5, 17, 1, 4, 99, 8, 25}
	fmt.Println(arr)
	BubbleSort(arr)
	fmt.Println(arr)

	BubbleSort(arr, func(a, b int) bool {
		return a < b
	})

	fmt.Println(arr)
}
