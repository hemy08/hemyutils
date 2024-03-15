package goutils

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

const (
	twoBytesLen   = 2
	fourBytesLen  = 4
	eightBytesLen = 8

	Ipv4BitLen = 32
	Ipv6BitLen = 128

	BitNumInAByte = 8

	DecimalBase     = 10
	HexadecimalBase = 16

	DefaultBitSizeForParseUint = 0

	InvalidSd = "FFFFFF" // 无效Sd
	BlankStr  = ""
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

// Uint16ToByteSlice 将uint16类型的值转换为长度为2的byte切片，大端序
func Uint16ToByteSlice(val uint16) []byte {
	bs := make([]byte, 2, 2)
	binary.BigEndian.PutUint16(bs, val)
	return bs
}

// Uint32ToByteSlice 将uint32类型的值转换为长度为4的byte切片，大端序
func Uint32ToByteSlice(val uint32) []byte {
	bs := make([]byte, 4, 4)
	binary.BigEndian.PutUint32(bs, val)
	return bs
}

// Int32ToByteSlice 将int32类型的值转换为长度为4的byte切片，大端序
func Int32ToByteSlice(val int32) ([]byte, error) {
	bs := bytes.NewBuffer([]byte{})
	if err := binary.Write(bs, binary.BigEndian, val); err != nil {
		return nil, fmt.Errorf("Int32ToByteSlice : Write failed [err:%v]", err)
	}
	return bs.Bytes(), nil
}

// Uint64ToByteSlice 将uint64类型的值转换为长度为8的byte切片，大端序
func Uint64ToByteSlice(val uint64) []byte {
	bs := make([]byte, 8, 8)
	binary.BigEndian.PutUint64(bs, val)
	return bs
}

// Int64ToByteSlice 将int64类型的值转换为长度为8的byte切片，大端序
func Int64ToByteSlice(val int64) ([]byte, error) {
	bs := bytes.NewBuffer([]byte{})
	if err := binary.Write(bs, binary.BigEndian, val); err != nil {
		return nil, fmt.Errorf("Int64ToByteSlice : Write failed [err:%v]", err)
	}
	return bs.Bytes(), nil
}

// StringToByteSlice 字符串转字节切片
func StringToByteSlice(str string) []byte {
	bs := bytes.NewBuffer([]byte{})
	bs.WriteString(str)
	return bs.Bytes()
}

// HexStringToByteSlice 将16进制字符串转换为byte切片
// 例如，string("0123456789abcdefABCDEF")转换为[]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xAB, 0xCD, 0xEF}
func HexStringToByteSlice(hexStr string) []byte {
	result, err := hex.DecodeString(hexStr)
	if err != nil {
		// logs.Error("DecodeString failed, err: %v", err)
		return nil
	}
	return result
}

// BytesToString Convert
func BytesToString(b []byte) string {
	var valuesText bytes.Buffer
	for i := range b {
		valuesText.WriteString(strconv.Itoa(int(b[i])))
	}
	return valuesText.String()
}

// Uint32ToDecimalString uint32转10进制字符串
func Uint32ToDecimalString(num uint32) string {
	return strconv.FormatUint(uint64(num), DecimalBase)
}

// Uint32ToHexString uint32转16进制字符串
func Uint32ToHexString(num uint32) string {
	return strconv.FormatUint(uint64(num), HexadecimalBase)
}

// Uint64ToDecimalString uint64转10进制字符串
func Uint64ToDecimalString(num uint64) string {
	return strconv.FormatUint(num, DecimalBase)
}

// Uint64ToHexString uint64转16进制字符串
func Uint64ToHexString(num uint64) string {
	return strconv.FormatUint(num, HexadecimalBase)
}

// Int64ToDecimalString int64转10进制字符串
func Int64ToDecimalString(num int64) string {
	return strconv.FormatInt(num, DecimalBase)
}

// ByteToHexString 例如，将byte :0xAB,转换为string("AB")
func ByteToHexString(bt byte) string {
	return fmt.Sprintf("%02X", bt)
}

// ByteToHexStringLower 例如，将byte :0xAB,转换为string("ab")
func ByteToHexStringLower(bt byte) string {
	return fmt.Sprintf("%02x", bt)
}

// ByteSliceToHexString 将byte切片转换为大写的16进制字符串
// 例如，将[]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}转换为string("0123456789ABCDEF")
func ByteSliceToHexString(bytes []byte) string {
	const hextable = "0123456789ABCDEF"

	dst := make([]byte, len(bytes)*twoBytesLen, len(bytes)*twoBytesLen)
	j := 0
	for _, v := range bytes {
		dst[j] = hextable[v>>fourBytesLen]
		dst[j+1] = hextable[v&0x0f]
		j += twoBytesLen
	}

	return string(dst)
}

// ByteSliceToHexStringLower 将byte切片转换为小写的16进制字符串
// 例如，将[]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}转换为string("0123456789abcdef")
func ByteSliceToHexStringLower(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

// ImmutableStringByteSliceToString 将字符串从其[]byte形式转换成string形式，相比于string([]byte)强转形式，性能高，无额外对象分配，前提：原[]byte不会被修改
// 注意使用场景：生命周期仅为方法内的[]byte，需要转换为string时，且该[]byte在转换后不会被修改，才能使用该方法转换
func ImmutableStringByteSliceToString(s []byte) string {
	return *(*string)(unsafe.Pointer(&s))
}

// Uint32ToHexEncodeString 无符号转换成16进制的字符串，且不带0x
// 例如，十进制1051转换成string("41D")
func Uint32ToHexEncodeString(teid uint32) string {
	return strings.ToUpper(hex.EncodeToString(Uint32ToByteSlice(teid)))
}

// TimeStampToString 时间转string
func TimeStampToString(TimeStamp int64) string {
	tmpTime := time.Unix(TimeStamp, 0).String()
	arr := strings.Fields(tmpTime)
	if len(arr) <= 1 {
		return ""
	}

	return arr[0] + " " + arr[1]
}

// Bool2String bool转string
func Bool2String(bool bool) string {
	if bool {
		return "true"
	}
	return "false"
}

// ByteSliceToUint16 将长度为1-2的byte切片转换为uint16类型的值，大端序
func ByteSliceToUint16(bytes []byte) uint16 {
	if len(bytes) > twoBytesLen || len(bytes) == 0 {
		return 0
	}
	var result uint16
	for idx, val := range bytes {
		result += uint16(val) << uint8((len(bytes)-1-idx)*BitNumInAByte)
	}
	return result
}

// ByteSliceToUint16Little 将长度为1-2的byte切片转换为uint16类型的值，小端序
func ByteSliceToUint16Little(bytes []byte) uint16 {
	if len(bytes) > twoBytesLen || len(bytes) == 0 {
		return 0
	}
	var result uint16
	for idx, val := range bytes {
		result += uint16(val) << (uint8(idx) * BitNumInAByte)
	}
	return result
}

// ByteSliceToUint32 将长度为1-4的byte切片转换为uint32类型的值，大端序
func ByteSliceToUint32(bytes []byte) uint32 {
	if len(bytes) > fourBytesLen || len(bytes) == 0 {
		return 0
	}
	var result uint32
	for idx, val := range bytes {
		result += uint32(val) << uint8((len(bytes)-1-idx)*BitNumInAByte)
	}
	return result
}

// ByteSliceToUint32Little 将长度为1-4的byte切片转换为uint32类型的值，小端序
func ByteSliceToUint32Little(bytes []byte) uint32 {
	if len(bytes) > fourBytesLen || len(bytes) == 0 {
		return 0
	}
	var result uint32
	for idx, val := range bytes {
		result += uint32(val) << (uint8(idx) * BitNumInAByte)
	}
	return result
}

// ByteSliceToUint64 将长度为1-8的byte切片转换为uint64类型的值，大端序
func ByteSliceToUint64(bytes []byte) uint64 {
	if len(bytes) > eightBytesLen || len(bytes) == 0 {
		return 0
	}
	var result uint64
	for idx, val := range bytes {
		result += uint64(val) << uint8((len(bytes)-1-idx)*BitNumInAByte)
	}
	return result
}

// HexStringToUint32 16进制字符串转换成整型
// 例如，string("0x111121")转换成1118497
func HexStringToUint32(str string) uint32 {
	str = strings.TrimPrefix(str, "0x")

	res, err := strconv.ParseInt(str, HexadecimalBase, DefaultBitSizeForParseUint)
	if err != nil {
		// logs.Info("strconv.ParseInt failed, err: %v", err)
		return 0
	}

	return uint32(res)
}
