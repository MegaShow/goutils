package ucrypto

import (
	"testing"

	"go.icytown.com/utils/internal/assert"
)

func TestMD5Hex(t *testing.T) {
	assert.Equal(t, "d41d8cd98f00b204e9800998ecf8427e", MD5Hex(""))
	assert.Equal(t, "e10adc3949ba59abbe56e057f20f883e", MD5Hex("123456"))
}

func TestSHA1Hex(t *testing.T) {
	assert.Equal(t, "da39a3ee5e6b4b0d3255bfef95601890afd80709", SHA1Hex(""))
	assert.Equal(t, "7c4a8d09ca3762af61e59520943dc26494f8941b", SHA1Hex("123456"))
}

func TestSHA256Hex(t *testing.T) {
	assert.Equal(t, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", SHA256Hex(""))
	assert.Equal(t, "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92", SHA256Hex("123456"))
}

func TestSHA512Hex(t *testing.T) {
	assert.Equal(t, "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e", SHA512Hex(""))
	assert.Equal(t, "ba3253876aed6bc22d4a6ff53d8406c6ad864195ed144ab5c87621b6c233b548baeae6956df346ec8c17f5ea10f35ee3cbc514797ed7ddd3145464e2a0bab413", SHA512Hex("123456"))
}
