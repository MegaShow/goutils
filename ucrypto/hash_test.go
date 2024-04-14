package ucrypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5(t *testing.T) {
	assert.Equal(t, "d41d8cd98f00b204e9800998ecf8427e", MD5(""))
	assert.Equal(t, "e10adc3949ba59abbe56e057f20f883e", MD5("123456"))
}

func TestSHA512(t *testing.T) {
	assert.Equal(t, "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e", SHA512(""))
	assert.Equal(t, "ba3253876aed6bc22d4a6ff53d8406c6ad864195ed144ab5c87621b6c233b548baeae6956df346ec8c17f5ea10f35ee3cbc514797ed7ddd3145464e2a0bab413", SHA512("123456"))
}
