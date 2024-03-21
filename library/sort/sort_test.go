package sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SelectSort_Case(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntSelectSort(intSlice, Swap)
	fmt.Printf("%v\n", intSlice)
	assert.Equal(t, intSlice, []int{13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	// intSlice = []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntSelectSort(intSlice, func(a, b int) bool { return a < b })
	fmt.Printf("%v\n", intSlice)
	assert.Equal(t, intSlice, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
}

func Test_BubbleSort_Case(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntBubbleSort(intSlice, func(a, b int) bool { return a < b })
	fmt.Printf("%v\n", intSlice)
	assert.Equal(t, intSlice, []int{13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	// intSlice = []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntBubbleSort(intSlice, func(a, b int) bool { return a > b })
	fmt.Printf("%v\n", intSlice)
	assert.Equal(t, intSlice, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
}

func Test_InsertSort_Case(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntInsertSort(intSlice, func(a, b int) bool { return a > b })
	fmt.Printf("%v\n", intSlice)
	assert.Equal(t, intSlice, []int{13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	// intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntInsertSort(intSlice, func(a, b int) bool { return a < b })
	fmt.Printf("%v\n", intSlice)
	assert.Equal(t, intSlice, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
}

func Test_ShellSort_Case(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntShellSort(intSlice, func(a, b int) bool { return a > b })
	fmt.Printf("%v\n", intSlice)
	assert.Equal(t, intSlice, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
	// intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntShellSort(intSlice, func(a, b int) bool { return a < b })
	fmt.Printf("%v\n", intSlice)
	assert.Equal(t, intSlice, []int{13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
}

func Test_MergeSort_Case(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntMergeSort(intSlice, 0, len(intSlice)-1, func(a, b int) bool { return a <= b })
	fmt.Printf("%v\n", intSlice)
	assert.Equal(t, intSlice, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
	// intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntMergeSort(intSlice, 0, len(intSlice)-1, func(a, b int) bool { return a > b })
	fmt.Printf("%v\n", intSlice)
	assert.Equal(t, intSlice, []int{13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
}

func Test_QuickSort_Case(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntQuickSort(intSlice, 0, len(intSlice)-1, func(a, b int) bool { return a <= b })
	fmt.Printf("%v\n", intSlice)
	assert.Equal(t, intSlice, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
	// intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntQuickSort(intSlice, 0, len(intSlice)-1, func(a, b int) bool { return a > b })
	fmt.Printf("%v\n", intSlice)
	assert.Equal(t, intSlice, []int{13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
}

func Test_HeapSort_Case(t *testing.T) {
	intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntHeapSort(intSlice, func(a, b int) bool { return a > b })
	fmt.Printf("%v\n", intSlice)
	assert.Equal(t, intSlice, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
	// intSlice := []int{9, 13, 11, 5, 2, 12, 6, 3, 1, 4, 8, 10, 7}
	IntHeapSort(intSlice, func(a, b int) bool { return a < b })
	fmt.Printf("%v\n", intSlice)
	assert.Equal(t, intSlice, []int{13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
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

func Test_ByteQuickSort_CaseUp(t *testing.T) {
	byteSlice := []byte("sdafoiajsgahfaneuasdfjhas")
	ByteQuickSort(byteSlice, 0, len(byteSlice)-1, func(a, b int) bool { return byteSlice[a] <= byteSlice[b] })
	fmt.Printf("%v\n", string(byteSlice))
	assert.Equal(t, "aaaaaaddefffghhijjnossssu", string(byteSlice))
	byteSlice = []byte("sdafoiajsgahfaneuasdfjhas")
	ByteQuickSort(byteSlice, 0, len(byteSlice)-1, func(a, b int) bool { return byteSlice[a] > byteSlice[b] })
	fmt.Printf("%v\n", string(byteSlice))
	assert.Equal(t, "ussssonjjihhgfffeddaaaaaa", string(byteSlice))
}
