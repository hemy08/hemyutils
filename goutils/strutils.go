package goutils

import (
	"sort"
	"strings"
	"unsafe"
)

// Str2Bytes convert string to array of byte
func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// Bytes2Str convert array of byte to string
func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// JointStr join str
func JointStr(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		_, err := sb.WriteString(str)
		if err != nil {
			return ""
		}
	}
	return sb.String()
}

// ClearStringMemory clear string memory
func ClearStringMemory(s string) {
	// 对单个字符表示的字符串，Go的runtime实现的字符串共享字符数据，位于统一的一个staticbytes数组中
	// 所以禁止清除它，否则程序后面运行时所有从单个字符解析出的字符串为'\0'值；避免设置口令、密码等敏感信息的长度小于等于1
	if len(s) <= 1 {
		return
	}

	bs := *(*[]byte)(unsafe.Pointer(&s))
	for i := 0; i < len(bs); i++ {
		bs[i] = 0
	}
}

func StringSliceEqual(leftSlice []string, rightSlice []string) bool {
	if len(leftSlice) != len(rightSlice) {
		return false
	}
	if (leftSlice == nil) != (rightSlice == nil) {
		return false
	}
	sort.Strings(leftSlice)
	sort.Strings(rightSlice)
	for idx, str := range leftSlice {
		if str != rightSlice[idx] {
			return false
		}
	}
	return true
}

// ClearByteArray clear byte array
func ClearByteArray(bt []byte) {
	btLen := len(bt)
	for i := 0; i < btLen; i++ {
		bt[i] = 0
	}
}
