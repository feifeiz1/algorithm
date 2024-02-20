package sort

import "cmp"

// SelectSort 选择排序，每一趟选出最大的或者最小的，放在前面
// 平均时间复杂度: O(n^2)
// 最坏时间复杂度: O(n^2)
// 最好时间复杂度: O(n^2)
// 空间复杂度: O(1)
// 是否稳定: 不稳定
func SelectSort[T cmp.Ordered](arr []T, o Ordered) {
	if len(arr) == 0 {
		return
	}
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if o == ASC {
				if arr[minIdx] > arr[j] {
					minIdx = j
				}
			} else if o == DESC {
				if arr[minIdx] < arr[j] {
					minIdx = j
				}
			}

		}
		arr[minIdx], arr[i] = arr[i], arr[minIdx]
	}
}
