package sort

import "sort"

func Swap(a, b int) bool { return a > b }

// IntSelectSort 选择排序
func IntSelectSort(arr []int, f func(a, b int) bool) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if f(arr[j], arr[minIndex]) {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// IntBubbleSort 冒泡排序
func IntBubbleSort(arr []int, f func(a, b int) bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if f(arr[j], arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// IntInsertSort 插入排序
func IntInsertSort(arr []int, f func(a, b int) bool) {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && f(arr[j], arr[j-1]); j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

// IntShellSort 希尔排序
func IntShellSort(arr []int, f func(a, b int) bool) {
	n := len(arr)
	gap := n / 2
	for gap > 0 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i
			for j >= gap && f(arr[j-gap], temp) {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
		gap /= 2
	}
}

func merge(arr []int, l, m, r int, f func(a, b int) bool) {
	n1 := m - l + 1
	n2 := r - m

	L := make([]int, n1)
	R := make([]int, n2)

	for i := 0; i < n1; i++ {
		L[i] = arr[l+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}

	i := 0
	j := 0
	k := l
	for i < n1 && j < n2 {
		if f(L[i], R[j]) {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}

// IntMergeSort 归并排序
func IntMergeSort(arr []int, l, r int, f func(a, b int) bool) {
	if l < r {
		m := (l + r) / 2
		IntMergeSort(arr, l, m, f)
		IntMergeSort(arr, m+1, r, f)
		merge(arr, l, m, r, f)
	}
}

// IntQuickSort 快速排序
func IntQuickSort(arr []int, left, right int, f func(a, b int) bool) {
	if left < right {
		pivot := partition(arr, left, right, f)
		IntQuickSort(arr, left, pivot-1, f)
		IntQuickSort(arr, pivot+1, right, f)
	}
}

func partition(arr []int, left, right int, f func(a, b int) bool) int {
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if f(arr[j], pivot) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

func maxHeapIfy(arr []int, n int, i int, f func(a, b int) bool) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && f(arr[l], arr[largest]) {
		largest = l
	}

	if r < n && f(arr[r], arr[largest]) {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeapIfy(arr, n, largest, f)
	}
}

// IntHeapSort 堆排序
func IntHeapSort(arr []int, f func(a, b int) bool) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		maxHeapIfy(arr, n, i, f)
	}

	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		maxHeapIfy(arr, i, 0, f)
	}
}

func bucketMaxMin(arr []int) (int, int) {
	// 找到最大值和最小值
	minVal, maxVal := arr[0], arr[0]
	for _, val := range arr {
		if val < minVal {
			minVal = val
		}
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal, minVal
}

func insertionSort(arr []int, f func(a, b int) bool) []int {
	var n = len(arr)
	for i := 1; i < n; i++ {
		var key = arr[i]
		var j = i - 1
		for j >= 0 && f(arr[j], key) {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	return arr
}

// FloatBucketSort 桶排序
func FloatBucketSort(arr []int, f func(a, b int) bool) {
	minVal, maxVal := bucketMaxMin(arr)

	// 计算桶的数量和桶的范围
	bucketSize := float64(maxVal-minVal+1) / float64(len(arr))
	bucketCount := int(bucketSize) + 1
	buckets := make([][]int, bucketCount)
	// 将数据分配到桶中

	for _, val := range arr {
		bucketIndex := int((float64(val) - float64(minVal)) / bucketSize)
		buckets[bucketIndex] = append(buckets[bucketIndex], val)
	}

	// 对每个桶进行排序
	for i := 0; i < bucketCount; i++ {
		sort.Ints(buckets[i])
	}

	// 合并桶中的数据
	index := 0
	for _, bucket := range buckets {
		for _, val := range bucket {
			arr[index] = val
			index++
		}
	}
}

// ByteQuickSort 快速排序
func ByteQuickSort(arr []byte, left, right int, f func(a, b int) bool) {
	if left < right {
		pivot := bytePartition(arr, left, right, f)
		ByteQuickSort(arr, left, pivot-1, f)
		ByteQuickSort(arr, pivot+1, right, f)
	}
}

func bytePartition(arr []byte, left, right int, f func(a, b int) bool) int {
	i := left
	for j := left; j < right; j++ {
		if f(j, right) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return i
}
