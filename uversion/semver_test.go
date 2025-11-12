package uversion

import (
	"encoding/json"
	"testing"

	"go.icytown.com/utils/internal/assert"
)

func TestNewSemVer(t *testing.T) {
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
}

func TestMustNewSemVer(t *testing.T) {
	assert.Equal(t, "1.0.0", MustNewSemVer(1, 0, 0, "", "").String())
	assert.Panics(t, func() { MustNewSemVer(1, 0, 0, "", "fd12+") })
}

func TestParseSemVer(t *testing.T) {
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
}

func TestMustParseSemVer(t *testing.T) {
	assert.Equal(t, "1.0.0", MustParseSemVer("1.0.0").String())
	assert.Panics(t, func() { MustParseSemVer("0.1.3+fd12+") })
}

func TestSemVer_Major(t *testing.T) {
	assert.Equal(t, uint64(1), MustParseSemVer("1.2.3").Major())
	assert.Equal(t, uint64(0), MustParseSemVer("0.1.4").Major())
}

func TestSemVer_Minor(t *testing.T) {
	assert.Equal(t, uint64(2), MustParseSemVer("1.2.3").Minor())
	assert.Equal(t, uint64(1), MustParseSemVer("0.1.4").Minor())
}

func TestSemVer_Patch(t *testing.T) {
	assert.Equal(t, uint64(3), MustParseSemVer("1.2.3").Patch())
	assert.Equal(t, uint64(4), MustParseSemVer("0.1.4").Patch())
}

func TestSemVer_PreRelease(t *testing.T) {
	assert.Equal(t, "", MustParseSemVer("1.2.3").PreRelease())
	assert.Equal(t, "alpha", MustParseSemVer("1.2.3-alpha+fd").PreRelease())
}

func TestSemVer_Build(t *testing.T) {
	assert.Equal(t, "", MustParseSemVer("1.2.3").Build())
	assert.Equal(t, "fd", MustParseSemVer("1.2.3-alpha+fd").Build())
}

func TestSemVer_String(t *testing.T) {
	assert.Equal(t, "1.2.3", MustParseSemVer("1.2.3").String())
	assert.Equal(t, "1.2.3-alpha+fd", MustParseSemVer("1.2.3-alpha+fd").String())
}

func TestSemVer_MarshalJSON(t *testing.T) {
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

func TestSemVer_UnmarshalJSON(t *testing.T) {
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
}

func TestSemVer_IncrMajor(t *testing.T) {
	v := MustParseSemVer("0.1.2")
	v.IncrMajor()
	assert.Equal(t, "1.0.0", v.String())

	v = MustParseSemVer("0.1.2-alpha")
	v.IncrMajor()
	assert.Equal(t, "1.0.0", v.String())
}

func TestSemVer_IncrMinor(t *testing.T) {
	v := MustParseSemVer("0.1.2")
	v.IncrMinor()
	assert.Equal(t, "0.2.0", v.String())

	v = MustParseSemVer("0.1.2-alpha")
	v.IncrMinor()
	assert.Equal(t, "0.2.0", v.String())
}

func TestSemVer_IncrPatch(t *testing.T) {
	v := MustParseSemVer("0.1.2")
	v.IncrPatch()
	assert.Equal(t, "0.1.3", v.String())

	v = MustParseSemVer("0.1.2-alpha")
	v.IncrPatch()
	assert.Equal(t, "0.1.3", v.String())
}

func TestSemVer_SetMajor(t *testing.T) {
	v := MustParseSemVer("0.1.2")
	v.SetMajor(10)
	assert.Equal(t, "10.1.2", v.String())
}

func TestSemVer_SetMinor(t *testing.T) {
	v := MustParseSemVer("0.1.2")
	v.SetMinor(5)
	assert.Equal(t, "0.5.2", v.String())
}

func TestSemVer_SetPatch(t *testing.T) {
	v := MustParseSemVer("0.1.2")
	v.SetPatch(10)
	assert.Equal(t, "0.1.10", v.String())
}

func TestSemVer_SetPreRelease(t *testing.T) {
	v := MustParseSemVer("0.1.2-alpha")
	err := v.SetPreRelease("beta.1")
	assert.Nil(t, err)
	assert.Equal(t, "0.1.2-beta.1", v.String())

	err = v.SetPreRelease("beta+1")
	assert.NotNil(t, err)

	err = v.SetPreRelease("")
	assert.Nil(t, err)
	assert.Equal(t, "0.1.2", v.String())
}

func TestSemVer_SetBuild(t *testing.T) {
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

func TestSemVer_Compare(t *testing.T) {
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
}
