package xstring

import "unsafe"

// BytesToString converts bytes to string by using package unsafe.
//
// 使用包 unsafe 将 bytes 转成 string.
func BytesToString(bs []byte) string {
	return unsafe.String(unsafe.SliceData(bs), len(bs)) //nolint:gosec // 零拷贝转换, 字符串使用期间不得修改底层字节
}

// StringToBytes converts string to bytes by using package unsafe.
//
// 使用包 unsafe 将 bytes 转成 string.
func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s)) //nolint:gosec // 零拷贝只读视图, 调用方不得修改返回的字节
}
