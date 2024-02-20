package sort

import "cmp"

// InsertSort 插入排序：从未排序序列中依次取数据，插入到已排序序列中相应的位置
func InsertSort[T cmp.Ordered](arr []T, o Ordered) {
	if len(arr) == 0 {
		return
	}
	for i := range arr {
		preIdx := i - 1
		cur := arr[i]
		if o == ASC {
			for preIdx >= 0 && arr[preIdx] < cur {
				arr[preIdx+1] = arr[preIdx]
				preIdx -= 1
			}
		} else if o == DESC {
			for preIdx >= 0 && arr[preIdx] > cur {
				arr[preIdx+1] = arr[preIdx]
				preIdx -= 1
			}
		}
		arr[preIdx+1] = cur
	}
}
