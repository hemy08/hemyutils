package shuffle

import (
	"math/rand"
	"time"
)

// IntSliceShuffle 洗牌算法
func IntSliceShuffle(arr []int) {
	rand.Seed(time.Now().UnixNano())
	var i, j int
	var temp int
	for i = len(arr) - 1; i > 0; i-- {
		j = rand.Intn(i + 1)
		temp = arr[i]
		arr[i] = arr[j]
		arr[j] = temp
	}
}

// StringSliceShuffle 洗牌算法
func StringSliceShuffle(arr []string) {
	rand.Seed(time.Now().UnixNano())
	var i, j int
	var temp string
	for i = len(arr) - 1; i > 0; i-- {
		j = rand.Intn(i + 1)
		temp = arr[i]
		arr[i] = arr[j]
		arr[j] = temp
	}
}
