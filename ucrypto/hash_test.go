package ucrypto

import "testing"

func TestSha512(t *testing.T) {
	tests := []struct {
		data string
		want string
	}{
		{data: "", want: "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"},
		{data: "123456", want: "ba3253876aed6bc22d4a6ff53d8406c6ad864195ed144ab5c87621b6c233b548baeae6956df346ec8c17f5ea10f35ee3cbc514797ed7ddd3145464e2a0bab413"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Sha512(tt.data); got != tt.want {
				t.Errorf("Sha512() = %v, want %v", got, tt.want)
			}
		})
	}
}
