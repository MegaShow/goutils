// Package ucrypto provides utils of crypto.
//
// 包 ucrypto 提供了密码学相关的工具.
package ucrypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"go.icytown.com/utils/internal/xstring"
)

// MD5Hex returns the MD5 checksum of the data as hex string.
//
// 以十六进制字符串形式返回 MD5 哈希值.
func MD5Hex(data string) string {
	sum := md5.Sum(xstring.StringToBytes(data))
	return hex.EncodeToString(sum[:])
}

// SHA1Hex returns the SHA1 checksum of the data as hex string.
//
// 以十六进制字符串形式返回 SHA1 哈希值.
func SHA1Hex(data string) string {
	sum := sha1.Sum(xstring.StringToBytes(data))
	return hex.EncodeToString(sum[:])
}

// SHA256Hex returns the SHA1 checksum of the data as hex string.
//
// 以十六进制字符串形式返回 SHA1 哈希值.
func SHA256Hex(data string) string {
	sum := sha256.Sum256(xstring.StringToBytes(data))
	return hex.EncodeToString(sum[:])
}

// SHA512Hex returns the SHA512 checksum of the data as hex string.
//
// 以十六进制字符串形式返回 SHA512 哈希值.
func SHA512Hex(data string) string {
	sum := sha512.Sum512(xstring.StringToBytes(data))
	return hex.EncodeToString(sum[:])
}
