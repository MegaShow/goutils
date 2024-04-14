// Package ucrypto provides utils of crypto.
//
// 包 ucrypto 提供了密码学相关的工具.
package ucrypto

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
)

// MD5 returns the MD5 checksum of the data as hex string.
//
// 以十六进制字符串形式返回 MD5 哈希值.
func MD5(data string) string {
	sum := md5.Sum([]byte(data))
	return hex.EncodeToString(sum[:])
}

// SHA512 returns the SHA512 checksum of the data as hex string.
//
// 以十六进制字符串形式返回 SHA512 哈希值.
func SHA512(data string) string {
	sum := sha512.Sum512([]byte(data))
	return hex.EncodeToString(sum[:])
}
