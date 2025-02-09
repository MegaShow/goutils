// Package uversion provides some version implementations.
package uversion

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"go.icytown.com/utils/internal/xstring"
	"go.icytown.com/utils/ucond"
)

// SemVer is Semantic Versioning Sepecification to manage software or library.
//
// 管理软件和库的语义化版本标准.
type SemVer struct {
	major uint64
	minor uint64
	patch uint64
	pre   string
	build string
}

// NewSemVer creates a SemVer.
//
// 创建一个 SemVer.
func NewSemVer(major, minor, patch uint64, pre, build string) (SemVer, error) {
	if !isValidSemVerIdentifier(pre) {
		return SemVer{}, errors.New("invalid prerelease")
	} else if !isValidSemVerIdentifier(build) {
		return SemVer{}, errors.New("invalid build")
	}

	return SemVer{
		major: major,
		minor: minor,
		patch: patch,
		pre:   pre,
		build: build,
	}, nil
}

// MustNewSemVer creates a SemVer, if SemVer is invalid then panic.
//
// 创建一个 SemVer, 如果版本是非法的则 panic.
func MustNewSemVer(major, minor, patch uint64, pre, build string) SemVer {
	ver, err := NewSemVer(major, minor, patch, pre, build)
	if err != nil {
		panic(err)
	}
	return ver
}

// ParseSemVer parses string to SemVer with format major.minor.patch-pre+build.
//
// 解析字符串成版本, 格式为 major.minor.patch-pre+build.
func ParseSemVer(s string) (SemVer, error) {
	parts := strings.SplitN(s, ".", 3)
	if len(parts) != 3 {
		return SemVer{}, errors.New("invalid semver")
	}

	i := 2
	for _, split := range [2]byte{'-', '+'} {
		if index := strings.IndexByte(parts[i], split); index != -1 {
			parts = append(parts, parts[i][index+1:])
			parts[i] = parts[i][:index]
			i = len(parts) - 1
		} else {
			parts = append(parts, "")
		}
	}

	major, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return SemVer{}, fmt.Errorf("invalid major: %w", err)
	}
	minor, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return SemVer{}, fmt.Errorf("invalid minor: %w", err)
	}
	patch, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return SemVer{}, fmt.Errorf("invalid patch: %w", err)
	}

	if !isValidSemVerIdentifier(parts[3]) {
		return SemVer{}, errors.New("invalid prerelease")
	} else if !isValidSemVerIdentifier(parts[4]) {
		return SemVer{}, errors.New("invalid build")
	}

	return SemVer{
		major: major,
		minor: minor,
		patch: patch,
		pre:   parts[3],
		build: parts[4],
	}, nil
}

// ParseSemVer parses string to SemVer with format major.minor.patch-pre+build.
// If string is invalid then panic.
//
// 解析字符串成版本, 格式为 major.minor.patch-pre+build.
// 如果字符串是非法的则 panic.
func MustParseSemVer(s string) SemVer {
	ver, err := ParseSemVer(s)
	if err != nil {
		panic(err)
	}
	return ver
}

// Major returns major version.
//
// 返回主版本号.
func (v SemVer) Major() uint64 {
	return v.major
}

// Minor returns minor version.
//
// 返回次版本号.
func (v SemVer) Minor() uint64 {
	return v.minor
}

// Patch returns patch version.
//
// 返回修订版本号.
func (v SemVer) Patch() uint64 {
	return v.patch
}

// PreRelease returns pre release version.
//
// 返回先行版本号.
func (v SemVer) PreRelease() string {
	return v.pre
}

// Build returns build info.
//
// 返回版本编译信息.
func (v SemVer) Build() string {
	return v.build
}

// String converts SemVer to string.
//
// 将版本转成字符串.
func (v SemVer) String() string {
	var b strings.Builder
	b.WriteString(strconv.FormatUint(v.major, 10))
	b.WriteByte('.')
	b.WriteString(strconv.FormatUint(v.minor, 10))
	b.WriteByte('.')
	b.WriteString(strconv.FormatUint(v.patch, 10))

	if v.pre != "" {
		b.WriteByte('-')
		b.WriteString(v.pre)
	}
	if v.build != "" {
		b.WriteByte('+')
		b.WriteString(v.build)
	}

	return b.String()
}

// MarshalJSON marshal SemVer into json bytes.
//
// 序列化版本.
func (v SemVer) MarshalJSON() ([]byte, error) {
	return []byte(`"` + v.String() + `"`), nil
}

// UnmarshalJSON unmarshal json bytes to SemVer.
//
// 反序列化版本.
func (v *SemVer) UnmarshalJSON(data []byte) error {
	s := xstring.BytesToString(data)
	if len(s) == 0 {
		return nil
	}
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return errors.New("invalid semver")
	}

	ver, err := ParseSemVer(s[1 : len(s)-1])
	if err != nil {
		return err
	}
	*v = ver
	return nil
}

// IncrMajor increases major version.
//
// 添加主版本号.
func (v *SemVer) IncrMajor() {
	v.major++
	v.minor = 0
	v.patch = 0
	v.pre = ""
	v.build = ""
}

// IncrMinor increases minor version.
//
// 添加次版本号.
func (v *SemVer) IncrMinor() {
	v.minor++
	v.patch = 0
	v.pre = ""
	v.build = ""
}

// IncrPatch increases patch version.
//
// 添加修订版本号.
func (v *SemVer) IncrPatch() {
	v.patch++
	v.pre = ""
	v.build = ""
}

// SetMajor modifies major version.
//
// 修改主版本号.
func (v *SemVer) SetMajor(major uint64) {
	v.major = major
}

// SetMinor modifies minor version.
//
// 修改次版本号.
func (v *SemVer) SetMinor(minor uint64) {
	v.minor = minor
}

// SetPatch modifies patch version.
//
// 修改修订版本号.
func (v *SemVer) SetPatch(patch uint64) {
	v.patch = patch
}

// SetPreRelease modifies pre release version.
//
// 修改先行版本号.
func (v *SemVer) SetPreRelease(pre string) error {
	if !isValidSemVerIdentifier(pre) {
		return errors.New("invalid prerelease")
	}
	v.pre = pre
	return nil
}

// SetBuild modifies build info.
//
// 修改版本编译信息.
func (v *SemVer) SetBuild(build string) error {
	if !isValidSemVerIdentifier(build) {
		return errors.New("invalid build")
	}
	v.build = build
	return nil
}

// Compare compares two SemVer. If v greater than other, returns 1; if v less than other, returns -1; if same, returns 0.
//
// 对比两个版本. 如果当前版本大于给定版本, 返回1; 如果当前版本小于给定版本, 返回-1; 如果相等, 返回0.
func (v SemVer) Compare(other SemVer) int {
	// 对比 major/minor/patch
	if v.major != other.major {
		return ucond.If(v.major > other.major, 1, -1)
	} else if v.minor != other.minor {
		return ucond.If(v.minor > other.minor, 1, -1)
	} else if v.patch != other.patch {
		return ucond.If(v.patch > other.patch, 1, -1)
	}

	// 对比 pre
	if v.pre == other.pre {
		return 0
	} else if v.pre == "" {
		return 1
	} else if other.pre == "" {
		return -1
	}

	vParts := strings.Split(v.pre, ".")
	oParts := strings.Split(other.pre, ".")
	for i := 0; i < len(vParts) && i < len(oParts); i++ {
		if vParts[i] == oParts[i] {
			continue
		}
		vNum, vErr := strconv.ParseUint(vParts[i], 10, 64)
		oNum, oErr := strconv.ParseUint(oParts[i], 10, 64)
		if vErr != nil || oErr != nil {
			return ucond.If(vParts[i] > oParts[i], 1, -1)
		}
		return ucond.If(vNum > oNum, 1, -1)
	}
	return ucond.If(len(vParts) > len(oParts), 1, -1)
}

var semVerIdentifierReg = regexp.MustCompile(`^[0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*$`)

func isValidSemVerIdentifier(s string) bool {
	return s == "" || semVerIdentifierReg.MatchString(s)
}
