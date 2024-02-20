package sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	fmt.Println(arr)
	InsertSort(arr, DESC)
	fmt.Println(arr)
}

func TestInsertSortAsc(t *testing.T) {
	fmt.Println(arr)
	InsertSort(arr, ASC)
	fmt.Println(arr)
}
