package uversion

import (
	"encoding/json"
	"testing"

	"go.icytown.com/utils/internal/assert"
)

func TestMustNewSemVer(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "1.0.0", MustNewSemVer(1, 0, 0, "", "").String())
	assert.Panics(t, func() { MustNewSemVer(1, 0, 0, "", "fd12+") })
}

func TestMustParseSemVer(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "1.0.0", MustParseSemVer("1.0.0").String())
	assert.Panics(t, func() { MustParseSemVer("0.1.3+fd12+") })
}

func TestNewSemVer(t *testing.T) {
	t.Parallel()
	ver, err := NewSemVer(1, 0, 0, "", "")
	assert.Nil(t, err)
	assert.Equal(t, "1.0.0", ver.String())

	ver, err = NewSemVer(0, 1, 3, "alpha", "fd12")
	assert.Nil(t, err)
	assert.Equal(t, "0.1.3-alpha+fd12", ver.String())

	_, err = NewSemVer(0, 1, 3, "alpha+", "fd")
	assert.NotNil(t, err)
	_, err = NewSemVer(0, 1, 3, "alpha", "fd12+")
	assert.NotNil(t, err)

	// 测试预发布数字标识符不允许前导零
	_, err = NewSemVer(1, 2, 3, "alpha.01", "")
	assert.NotNil(t, err)
}

func TestParseSemVer(t *testing.T) {
	t.Parallel()
	ver, err := ParseSemVer("1.0.0")
	assert.Nil(t, err)
	assert.Equal(t, "1.0.0", ver.String())

	ver, err = ParseSemVer("0.1.3-beta")
	assert.Nil(t, err)
	assert.Equal(t, "0.1.3-beta", ver.String())

	ver, err = ParseSemVer("0.1.3-alpha+fd12")
	assert.Nil(t, err)
	assert.Equal(t, "0.1.3-alpha+fd12", ver.String())

	ver, err = ParseSemVer("0.1.3+fd12")
	assert.Nil(t, err)
	assert.Equal(t, "0.1.3+fd12", ver.String())

	_, err = ParseSemVer("1.0")
	assert.NotNil(t, err)
	_, err = ParseSemVer("x.0.0")
	assert.NotNil(t, err)
	_, err = ParseSemVer("1.y.0")
	assert.NotNil(t, err)
	_, err = ParseSemVer("1.0.z")
	assert.NotNil(t, err)
	_, err = ParseSemVer("0.1.3-alpha*")
	assert.NotNil(t, err)
	_, err = ParseSemVer("0.1.3-alpha+fd12+")
	assert.NotNil(t, err)

	// 测试构建标识符中的连字符和前导零
	for _, valid := range []string{"1.2.3+build-1", "1.2.3-alpha-1+build.01", "1.2.3-0"} {
		ver, err := ParseSemVer(valid)
		assert.Nil(t, err)
		assert.Equal(t, valid, ver.String())
	}

	// 测试前导零和空标识符
	for _, invalid := range []string{"01.2.3", "1.02.3", "1.2.03", "1.2.3-", "1.2.3+", "1.2.3-01", "1.2.3-alpha.01", "1.2.3-alpha..1"} {
		_, err := ParseSemVer(invalid)
		assert.NotNil(t, err)
	}
}

func TestSemVer_Build(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "", MustParseSemVer("1.2.3").Build())
	assert.Equal(t, "fd", MustParseSemVer("1.2.3-alpha+fd").Build())
}

func TestSemVer_Compare(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 0, MustParseSemVer("1.0.0").Compare(MustParseSemVer("1.0.0")))
	assert.Equal(t, -1, MustParseSemVer("1.0.0").Compare(MustParseSemVer("2.0.0")))
	assert.Equal(t, 1, MustParseSemVer("2.0.0").Compare(MustParseSemVer("1.0.0")))
	assert.Equal(t, -1, MustParseSemVer("2.0.0").Compare(MustParseSemVer("2.1.0")))
	assert.Equal(t, 1, MustParseSemVer("2.1.1").Compare(MustParseSemVer("2.1.0")))

	assert.Equal(t, -1, MustParseSemVer("1.0.0-alpha").Compare(MustParseSemVer("1.0.0")))
	assert.Equal(t, 1, MustParseSemVer("1.0.0").Compare(MustParseSemVer("1.0.0-alpha")))
	assert.Equal(t, 0, MustParseSemVer("1.0.0+fd").Compare(MustParseSemVer("1.0.0+ff")))

	assert.Equal(t, -1, MustParseSemVer("1.0.0-alpha").Compare(MustParseSemVer("1.0.0-alpha.1")))
	assert.Equal(t, -1, MustParseSemVer("1.0.0-alpha.1").Compare(MustParseSemVer("1.0.0-alpha.beta")))
	assert.Equal(t, -1, MustParseSemVer("1.0.0-alpha.beta").Compare(MustParseSemVer("1.0.0-beta")))
	assert.Equal(t, -1, MustParseSemVer("1.0.0-beta").Compare(MustParseSemVer("1.0.0-beta.2")))
	assert.Equal(t, -1, MustParseSemVer("1.0.0-beta.2").Compare(MustParseSemVer("1.0.0-beta.11")))
	assert.Equal(t, -1, MustParseSemVer("1.0.0-beta.11").Compare(MustParseSemVer("1.0.0-rc.1")))
	assert.Equal(t, -1, MustParseSemVer("1.0.0-rc.1").Compare(MustParseSemVer("1.0.0")))

	// 测试数字标识符低于非数字标识符, 且不受 uint64 上限影响
	for _, pair := range [][2]string{
		{"1.0.0-2", "1.0.0-1a"},
		{"1.0.0-18446744073709551615", "1.0.0-18446744073709551616"},
		{"1.0.0-99999999999999999999", "1.0.0-100000000000000000000"},
	} {
		left, right := MustParseSemVer(pair[0]), MustParseSemVer(pair[1])
		assert.Equal(t, -1, left.Compare(right))
		assert.Equal(t, 1, right.Compare(left))
	}
}

func TestSemVer_IncrMajor(t *testing.T) {
	t.Parallel()
	v := MustParseSemVer("0.1.2")
	v.IncrMajor()
	assert.Equal(t, "1.0.0", v.String())

	v = MustParseSemVer("0.1.2-alpha")
	v.IncrMajor()
	assert.Equal(t, "1.0.0", v.String())
}

func TestSemVer_IncrMinor(t *testing.T) {
	t.Parallel()
	v := MustParseSemVer("0.1.2")
	v.IncrMinor()
	assert.Equal(t, "0.2.0", v.String())

	v = MustParseSemVer("0.1.2-alpha")
	v.IncrMinor()
	assert.Equal(t, "0.2.0", v.String())
}

func TestSemVer_IncrPatch(t *testing.T) {
	t.Parallel()
	v := MustParseSemVer("0.1.2")
	v.IncrPatch()
	assert.Equal(t, "0.1.3", v.String())

	v = MustParseSemVer("0.1.2-alpha")
	v.IncrPatch()
	assert.Equal(t, "0.1.3", v.String())
}

func TestSemVer_Major(t *testing.T) {
	t.Parallel()
	assert.Equal(t, uint64(1), MustParseSemVer("1.2.3").Major())
	assert.Equal(t, uint64(0), MustParseSemVer("0.1.4").Major())
}

func TestSemVer_MarshalJSON(t *testing.T) {
	t.Parallel()
	ver, _ := ParseSemVer("1.2.3")
	bytes, err := json.Marshal(&ver)
	assert.Nil(t, err)
	assert.Equal(t, `"1.2.3"`, string(bytes))

	obj := struct {
		V SemVer `json:"v"`
	}{V: ver}
	bytes, err = json.Marshal(obj)
	assert.Nil(t, err)
	assert.Equal(t, `{"v":"1.2.3"}`, string(bytes))
}

func TestSemVer_Minor(t *testing.T) {
	t.Parallel()
	assert.Equal(t, uint64(2), MustParseSemVer("1.2.3").Minor())
	assert.Equal(t, uint64(1), MustParseSemVer("0.1.4").Minor())
}

func TestSemVer_Patch(t *testing.T) {
	t.Parallel()
	assert.Equal(t, uint64(3), MustParseSemVer("1.2.3").Patch())
	assert.Equal(t, uint64(4), MustParseSemVer("0.1.4").Patch())
}

func TestSemVer_PreRelease(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "", MustParseSemVer("1.2.3").PreRelease())
	assert.Equal(t, "alpha", MustParseSemVer("1.2.3-alpha+fd").PreRelease())
}

func TestSemVer_SetBuild(t *testing.T) {
	t.Parallel()
	v := MustParseSemVer("0.1.2-alpha")
	err := v.SetBuild("2025")
	assert.Nil(t, err)
	assert.Equal(t, "0.1.2-alpha+2025", v.String())

	err = v.SetBuild("2024+")
	assert.NotNil(t, err)

	err = v.SetBuild("")
	assert.Nil(t, err)
	assert.Equal(t, "0.1.2-alpha", v.String())
}

func TestSemVer_SetMajor(t *testing.T) {
	t.Parallel()
	v := MustParseSemVer("0.1.2")
	v.SetMajor(10)
	assert.Equal(t, "10.1.2", v.String())
}

func TestSemVer_SetMinor(t *testing.T) {
	t.Parallel()
	v := MustParseSemVer("0.1.2")
	v.SetMinor(5)
	assert.Equal(t, "0.5.2", v.String())
}

func TestSemVer_SetPatch(t *testing.T) {
	t.Parallel()
	v := MustParseSemVer("0.1.2")
	v.SetPatch(10)
	assert.Equal(t, "0.1.10", v.String())
}

func TestSemVer_SetPreRelease(t *testing.T) {
	t.Parallel()
	v := MustParseSemVer("0.1.2-alpha")
	err := v.SetPreRelease("beta.1")
	assert.Nil(t, err)
	assert.Equal(t, "0.1.2-beta.1", v.String())

	err = v.SetPreRelease("beta+1")
	assert.NotNil(t, err)

	err = v.SetPreRelease("")
	assert.Nil(t, err)
	assert.Equal(t, "0.1.2", v.String())

	// 测试非法预发布标识符不会覆盖原值
	err = v.SetPreRelease("alpha.01")
	assert.NotNil(t, err)
	assert.Equal(t, "", v.PreRelease())
}

func TestSemVer_String(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "1.2.3", MustParseSemVer("1.2.3").String())
	assert.Equal(t, "1.2.3-alpha+fd", MustParseSemVer("1.2.3-alpha+fd").String())
}

func TestSemVer_UnmarshalJSON(t *testing.T) {
	t.Parallel()
	var ver SemVer
	err := json.Unmarshal([]byte(`"1.2.3"`), &ver)
	assert.Nil(t, err)
	assert.Equal(t, "1.2.3", ver.String())

	obj1 := struct {
		V SemVer `json:"v"`
	}{}
	err = json.Unmarshal([]byte(`{"v":"1.2.3-alpha"}`), &obj1)
	assert.Nil(t, err)
	assert.Equal(t, "1.2.3-alpha", obj1.V.String())

	obj2 := struct {
		V *SemVer `json:"v"`
	}{}
	err = json.Unmarshal([]byte(`{"v":"1.2.3+fd"}`), &obj2)
	assert.Nil(t, err)
	assert.Equal(t, "1.2.3+fd", obj2.V.String())
	err = json.Unmarshal([]byte(`{"v":null}`), &obj2)
	assert.Nil(t, err)
	assert.Nil(t, obj2.V)

	// 测试合法 JSON 转义、null 和失败时保留原值
	err = json.Unmarshal([]byte(`"1.2.\u0033"`), &ver)
	assert.Nil(t, err)
	assert.Equal(t, "1.2.3", ver.String())
	err = json.Unmarshal([]byte(`null`), &ver)
	assert.Nil(t, err)
	assert.Equal(t, "1.2.3", ver.String())
	for _, invalid := range []string{`123`, `"invalid"`, `""`, `"1.2.3`, `"1.2.3\n"`} {
		err = json.Unmarshal([]byte(invalid), &ver)
		assert.NotNil(t, err)
		assert.Equal(t, "1.2.3", ver.String())
	}
	err = ver.UnmarshalJSON(nil)
	assert.NotNil(t, err)
}
