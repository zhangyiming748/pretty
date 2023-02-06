package arrutil_test

import (
	"fmt"
	"testing"

	"github.com/zhangyiming748/pretty/arrutil"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

func TestReverse(t *testing.T) {
	ss := []string{"a", "b", "c"}

	arrutil.Reverse(ss)
	assert.Eq(t, []string{"c", "b", "a"}, ss)
}

func TestStringsRemove(t *testing.T) {
	ss := []string{"a", "b", "c"}
	ns := arrutil.StringsRemove(ss, "b")

	assert.Contains(t, ns, "a")
	assert.NotContains(t, ns, "b")
	assert.Len(t, ns, 2)
}

func TestStringsFilter(t *testing.T) {
	is := assert.New(t)

	ss := arrutil.StringsFilter([]string{"a", "", "b", ""})
	is.Eq([]string{"a", "b"}, ss)
}

func TestTrimStrings(t *testing.T) {
	is := assert.New(t)

	// TrimStrings
	ss := arrutil.TrimStrings([]string{" a", "b ", " c "})
	is.Eq("[a b c]", fmt.Sprint(ss))
	ss = arrutil.TrimStrings([]string{",a", "b.", ",.c,"}, ",.")
	is.Eq("[a b c]", fmt.Sprint(ss))
	ss = arrutil.TrimStrings([]string{",a", "b.", ",.c,"}, ",", ".")
	is.Eq("[a b c]", fmt.Sprint(ss))
}

func TestGetRandomOne(t *testing.T) {
	is := assert.New(t)
	// int slice
	intSlice := []int{1, 2, 3, 4, 5, 6}
	intVal := arrutil.GetRandomOne(intSlice)
	intVal1 := arrutil.GetRandomOne(intSlice)
	for intVal == intVal1 {
		intVal1 = arrutil.GetRandomOne(intSlice)
	}

	assert.IsType(t, 0, intVal)
	is.True(arrutil.HasValue(intSlice, intVal))
	assert.IsType(t, 0, intVal1)
	is.True(arrutil.HasValue(intSlice, intVal1))
	assert.NotEq(t, intVal, intVal1)

	// int array
	intArray := []int{1, 2, 3, 4, 5, 6}
	intReturned := arrutil.GetRandomOne(intArray)
	intReturned1 := arrutil.GetRandomOne(intArray)
	for intReturned == intReturned1 {
		intReturned1 = arrutil.GetRandomOne(intArray)
	}
	assert.IsType(t, 0, intReturned)
	is.True(arrutil.Contains(intArray, intReturned))
	assert.IsType(t, 0, intReturned1)
	is.True(arrutil.Contains(intArray, intReturned1))
	assert.NotEq(t, intReturned, intReturned1)

	// string slice
	strSlice := []string{"aa", "bb", "cc", "dd"}
	strVal := arrutil.GetRandomOne(strSlice)
	strVal1 := arrutil.GetRandomOne(strSlice)
	for strVal == strVal1 {
		strVal1 = arrutil.GetRandomOne(strSlice)
	}

	assert.IsType(t, "", strVal)
	is.True(arrutil.Contains(strSlice, strVal))
	assert.IsType(t, "", strVal1)
	is.True(arrutil.Contains(strSlice, strVal1))
	assert.NotEq(t, strVal, strVal1)

	// string array
	strArray := []string{"aa", "bb", "cc", "dd"}
	strReturned := arrutil.GetRandomOne(strArray)
	strReturned1 := arrutil.GetRandomOne(strArray)
	for strReturned == strReturned1 {
		strReturned1 = arrutil.GetRandomOne(strArray)
	}

	assert.IsType(t, "", strReturned)
	is.True(arrutil.Contains(strArray, strReturned))
	assert.IsType(t, "", strReturned1)
	is.True(arrutil.Contains(strArray, strReturned1))
	assert.NotEq(t, strReturned, strReturned1)

	// byte slice
	byteSlice := []byte("abcdefg")
	byteVal := arrutil.GetRandomOne(byteSlice)
	byteVal1 := arrutil.GetRandomOne(byteSlice)
	for byteVal == byteVal1 {
		byteVal1 = arrutil.GetRandomOne(byteSlice)
	}

	assert.IsType(t, byte('a'), byteVal)
	is.True(arrutil.Contains(byteSlice, byteVal))
	assert.IsType(t, byte('a'), byteVal1)
	is.True(arrutil.Contains(byteSlice, byteVal1))
	assert.NotEq(t, byteVal, byteVal1)

	is.Panics(func() {
		arrutil.RandomOne([]int{})
	})
}

func TestUnique(t *testing.T) {
	assert.Eq(t, []int{2, 3, 4}, arrutil.Unique[int]([]int{2, 3, 2, 4}))
	assert.Eq(t, []uint{2, 3, 4}, arrutil.Unique([]uint{2, 3, 2, 4}))
	assert.Eq(t, []string{"ab", "bc", "cd"}, arrutil.Unique([]string{"ab", "bc", "ab", "cd"}))
}
