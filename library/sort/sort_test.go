package sort

import (
	"fmt"
	"testing"
)

func Test_SelectSort_CaseUp(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntSelectSort(intSlice, Swap)
	fmt.Printf("%v\n", intSlice)
}

func Test_SelectSort_CaseDown(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntSelectSort(intSlice, func(a, b int) bool { return a < b })
	fmt.Printf("%v\n", intSlice)
}

func Test_BubbleSort_CaseUp(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntBubbleSort(intSlice, Swap)
	fmt.Printf("%v\n", intSlice)
}

func Test_BubbleSort_CaseDown(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntBubbleSort(intSlice, func(a, b int) bool { return a < b })
	fmt.Printf("%v\n", intSlice)
}

func Test_InsertSort_CaseUp(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntInsertSort(intSlice, func(a, b int) bool { return a < b })
	fmt.Printf("%v\n", intSlice)
}

func Test_InsertSort_CaseDown(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntInsertSort(intSlice, func(a, b int) bool { return a > b })
	fmt.Printf("%v\n", intSlice)
}

func Test_ShellSort_CaseUp(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntShellSort(intSlice, func(a, b int) bool { return a > b })
	fmt.Printf("%v\n", intSlice)
}

func Test_ShellSort_CaseDown(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntShellSort(intSlice, func(a, b int) bool { return a > b })
	fmt.Printf("%v\n", intSlice)
}

func Test_MergeSort_CaseUp(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntMergeSort(intSlice, 0, len(intSlice)-1, func(a, b int) bool { return a <= b })
	fmt.Printf("%v\n", intSlice)
}

func Test_MergeSort_CaseDown(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntMergeSort(intSlice, 0, len(intSlice)-1, func(a, b int) bool { return a > b })
	fmt.Printf("%v\n", intSlice)
}

func Test_QuickSort_CaseUp(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntQuickSort(intSlice, 0, len(intSlice)-1, func(a, b int) bool { return a <= b })
	fmt.Printf("%v\n", intSlice)
}

func Test_QuickSort_CaseDown(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntQuickSort(intSlice, 0, len(intSlice)-1, func(a, b int) bool { return a > b })
	fmt.Printf("%v\n", intSlice)
}

func Test_HeapSort_CaseUp(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntHeapSort(intSlice, func(a, b int) bool { return a > b })
	fmt.Printf("%v\n", intSlice)
}

func Test_HeapSort_CaseDown(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntHeapSort(intSlice, func(a, b int) bool { return a < b })
	fmt.Printf("%v\n", intSlice)
}

func Test_BucketSort_CaseUp(t *testing.T) {
	intSlice := []int{9.0, 13.0, 11.0, 5.0, 2.0, 12.0, 6.0, 3.0, 1.0, 4.0, 8.0, 10.0, 7.0}
	FloatBucketSort(intSlice, func(a, b int) bool { return a > b })
	fmt.Printf("%v\n", intSlice)
}

func Test_BucketSort_CaseDown(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	FloatBucketSort(intSlice, func(a, b int) bool { return a < b })
	fmt.Printf("%v\n", intSlice)
}
