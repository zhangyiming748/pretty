package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestIsAlphabet(t *testing.T) {
	assert.True(t, strutil.IsNumChar('9'))
	assert.False(t, strutil.IsNumChar('A'))

	assert.False(t, strutil.IsAlphabet('9'))
	assert.False(t, strutil.IsAlphabet('+'))

	assert.True(t, strutil.IsAlphabet('A'))
	assert.True(t, strutil.IsAlphabet('a'))
	assert.True(t, strutil.IsAlphabet('Z'))
	assert.True(t, strutil.IsAlphabet('z'))

	assert.True(t, strutil.IsNumeric("234"))
	assert.False(t, strutil.IsNumeric("a34"))
}

func TestIsAlphaNum(t *testing.T) {
	assert.False(t, strutil.IsAlphaNum('+'))

	assert.True(t, strutil.IsAlphaNum('9'))
	assert.True(t, strutil.IsAlphaNum('A'))
	assert.True(t, strutil.IsAlphaNum('a'))
	assert.True(t, strutil.IsAlphaNum('Z'))
	assert.True(t, strutil.IsAlphaNum('z'))
}

func TestNoCaseEq(t *testing.T) {
	assert.True(t, strutil.Equal("a", "a"))
	assert.True(t, strutil.NoCaseEq("A", "a"))
	assert.True(t, strutil.NoCaseEq("Ab", "aB"))
	assert.False(t, strutil.Equal("a", "b"))
}

func TestStrPos(t *testing.T) {
	// StrPos
	assert.Eq(t, -1, strutil.StrPos("xyz", "a"))
	assert.Eq(t, 0, strutil.StrPos("xyz", "x"))
	assert.Eq(t, 2, strutil.StrPos("xyz", "z"))

	// RunePos
	assert.Eq(t, -1, strutil.RunePos("xyz", 'a'))
	assert.Eq(t, 0, strutil.RunePos("xyz", 'x'))
	assert.Eq(t, 2, strutil.RunePos("xyz", 'z'))
	assert.Eq(t, 5, strutil.RunePos("hi时间", '间'))

	// BytePos
	assert.Eq(t, -1, strutil.BytePos("xyz", 'a'))
	assert.Eq(t, 0, strutil.BytePos("xyz", 'x'))
	assert.Eq(t, 2, strutil.BytePos("xyz", 'z'))
	// assert.Eq(t, 2, strutil.BytePos("hi时间", '间')) // will build error
}

func TestIsStartOf(t *testing.T) {
	tests := []struct {
		give string
		sub  string
		want bool
	}{
		{"abc", "a", true},
		{"abc", "d", false},
	}

	for _, item := range tests {
		assert.Eq(t, item.want, strutil.HasPrefix(item.give, item.sub))
		assert.Eq(t, item.want, strutil.IsStartOf(item.give, item.sub))
	}

	assert.True(t, strutil.IsStartsOf("abc", []string{"a", "b"}))
	assert.False(t, strutil.IsStartsOf("abc", []string{"d", "e"}))
}

func TestIsEndOf(t *testing.T) {
	tests := []struct {
		give string
		sub  string
		want bool
	}{
		{"abc", "c", true},
		{"abc", "d", false},
		{"some.json", ".json", true},
	}

	for _, item := range tests {
		assert.Eq(t, item.want, strutil.HasSuffix(item.give, item.sub))
		assert.Eq(t, item.want, strutil.IsEndOf(item.give, item.sub))
	}
}

func TestIsSpace(t *testing.T) {
	assert.True(t, strutil.IsSpace(' '))
	assert.True(t, strutil.IsSpace('\n'))
	assert.True(t, strutil.IsSpaceRune('\n'))
	assert.True(t, strutil.IsSpaceRune('\t'))

	assert.False(t, strutil.IsBlank(" a "))
	assert.True(t, strutil.IsNotBlank(" a "))
	assert.False(t, strutil.IsEmpty(" "))
	assert.True(t, strutil.IsBlank(""))
	assert.True(t, strutil.IsBlank(" "))
	assert.True(t, strutil.IsBlank("   "))
	assert.False(t, strutil.IsNotBlank("   "))

	assert.False(t, strutil.IsBlankBytes([]byte(" a ")))
	assert.True(t, strutil.IsBlankBytes([]byte(" ")))
	assert.True(t, strutil.IsBlankBytes([]byte("   ")))
}

func TestIsSymbol(t *testing.T) {
	assert.False(t, strutil.IsSymbol('a'))
	assert.True(t, strutil.IsSymbol('●'))
}

func TestIsVersion(t *testing.T) {
	assert.False(t, strutil.IsVersion("abc"))
	assert.False(t, strutil.IsVersion(".2"))
	assert.False(t, strutil.IsVersion("a.2"))

	assert.True(t, strutil.IsVersion("0.1"))
	assert.True(t, strutil.IsVersion("0.1.0"))
	assert.True(t, strutil.IsVersion("1.2.0"))
	assert.True(t, strutil.IsVersion("1.2.0-beta"))
	assert.True(t, strutil.IsVersion("1.2.0-beta2"))
	assert.True(t, strutil.IsVersion("1.2.0-alpha1"))
}

func TestIEqual(t *testing.T) {
	assert.False(t, strutil.IEqual("h3ab2c", "d"))
	assert.False(t, strutil.IEqual("ab", "ac"))
	assert.True(t, strutil.IEqual("ab", "AB"))
	assert.True(t, strutil.IEqual("ab", "Ab"))
	assert.True(t, strutil.IEqual("ab", "ab"))
}

func TestIContains(t *testing.T) {
	assert.False(t, strutil.IContains("h3ab2c", "d"))
	assert.True(t, strutil.IContains("h3ab2c", "AB"))
	assert.True(t, strutil.IContains("H3AB2C", "aB"))
}

func TestHasOneSub(t *testing.T) {
	assert.False(t, strutil.ContainsOne("h3ab2c", []string{"d"}))
	assert.False(t, strutil.HasOneSub("h3ab2c", []string{"d"}))
	assert.True(t, strutil.HasOneSub("h3ab2c", []string{"ab"}))
}

func TestHasAllSubs(t *testing.T) {
	assert.False(t, strutil.HasAllSubs("h3ab2c", []string{"a", "d"}))
	assert.True(t, strutil.HasAllSubs("h3ab2c", []string{"a", "b"}))
	assert.True(t, strutil.ContainsAll("h3ab2c", []string{"a", "b"}))
}

func TestVersionCompare(t *testing.T) {
	versions := []struct{ a, b string }{
		{"1.0.221.9289", "1.05.00.0156"},
		// Go versions
		{"1", "1.0.1"},
		{"1.0.1", "1.0.2"},
		{"1.0.2", "1.0.3"},
		{"1.0.3", "1.1"},
		{"1.1", "1.1.1"},
		{"1.1.1", "1.1.2"},
		{"1.1.2", "1.2"},
	}
	for _, version := range versions {
		assert.True(t, strutil.VersionCompare(version.a, version.b, "<"), version.a+"<"+version.b)
		assert.True(t, strutil.VersionCompare(version.a, version.b, "<="), version.a+"<="+version.b)
		assert.True(t, strutil.VersionCompare(version.b, version.a, ">"), version.a+">"+version.b)
		assert.True(t, strutil.VersionCompare(version.b, version.a, ">="), version.a+">="+version.b)
	}

	assert.True(t, strutil.VersionCompare("1.0", "1.0", ""))
	assert.True(t, strutil.VersionCompare("1.0", "1.0", "="))

	assert.False(t, strutil.Compare("2020-12-16", "2021-12-17", ">="))
}
