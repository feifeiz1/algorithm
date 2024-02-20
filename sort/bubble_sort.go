package sort

import "cmp"

type cmpFunc[T cmp.Ordered] func(a, b T) bool

func ascCmpFunc[T cmp.Ordered](a, b T) bool {
	return a > b
}

// BubbleSort 冒泡排序，每轮遇到比自己大/小的就交换顺序
// 平均时间复杂度: O(n^2)
// 最好时间复杂度: O(n^2)
// 最坏时间复杂度: O(n)
// 空间复杂度: O(1)
// 是否稳定: 是
func BubbleSort[T cmp.Ordered](arr []T, f ...cmpFunc[T]) {
	if len(arr) <= 0 {
		return
	}
	cmpF := ascCmpFunc[T]
	if len(f) != 0 {
		cmpF = f[0]
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if cmpF(arr[i], arr[j]) {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
