package sort

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

func TestSelectSort(t *testing.T) {
	fmt.Println(arr)

	SelectSort(arr, ASC)
	fmt.Println(arr)

	SelectSort(arr, DESC)
	fmt.Println(arr)

	a := make([]float64, 10)
	for i := 0; i < 10; i++ {
		x := rand.Float64()
		a[i] = x
	}
	fmt.Println(a)

	SelectSort(a, ASC)
	fmt.Println(a)
}
