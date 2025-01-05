package xstring

import "unsafe"

// BytesToString converts bytes to string by using package unsafe.
//
// 使用包 unsafe 将 bytes 转成 string.
func BytesToString(bs []byte) string {
	return unsafe.String(unsafe.SliceData(bs), len(bs))
}

// StringToBytes converts string to bytes by using package unsafe.
//
// 使用包 unsafe 将 bytes 转成 string.
func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
