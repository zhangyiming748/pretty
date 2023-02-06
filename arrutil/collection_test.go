package arrutil_test

import (
	"testing"

	"github.com/zhangyiming748/pretty/arrutil"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

// StringEqualComparer tests
func TestStringEqualComparerShouldEquals(t *testing.T) {
	assert.Eq(t, 0, arrutil.StringEqualsComparer("a", "a"))
}

func TestStringEqualComparerShouldNotEquals(t *testing.T) {
	assert.NotEq(t, 0, arrutil.StringEqualsComparer("a", "b"))
}

func TestStringEqualComparerElementNotString(t *testing.T) {
	assert.Eq(t, -1, arrutil.StringEqualsComparer(1, "a"))
}

func TestStringEqualComparerPtr(t *testing.T) {
	ptrVal := "a"
	assert.Eq(t, 0, arrutil.StringEqualsComparer(&ptrVal, "a"))
}

// ReferenceEqualsComparer tests
func TestReferenceEqualsComparerShouldEquals(t *testing.T) {
	assert.Eq(t, 0, arrutil.ReferenceEqualsComparer(1, 1))
}

func TestReferenceEqualsComparerShouldNotEquals(t *testing.T) {
	assert.NotEq(t, 0, arrutil.ReferenceEqualsComparer(1, 2))
}

// ElemTypeEqualCompareFunc
func TestElemTypeEqualCompareFuncShouldEquals(t *testing.T) {
	assert.Eq(t, 0, arrutil.ElemTypeEqualsComparer(1, 2))
}

func TestElemTypeEqualCompareFuncShouldNotEquals(t *testing.T) {
	assert.NotEq(t, 0, arrutil.ElemTypeEqualsComparer(1, "2"))
}

func TestExceptsShouldPassed(t *testing.T) {
	data := []string{
		"a",
		"b",
		"c",
	}
	result := arrutil.Excepts(data, []string{"a", "b"}, arrutil.StringEqualsComparer)
	assert.Eq(t, []string{"c"}, result.([]string))
}

func TestExceptsFirstNotSliceShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			return
		} else {
			t.Fail()
		}
	}()
	arrutil.Excepts([1]string{"a"}, []string{"a", "b"}, arrutil.StringEqualsComparer)
}

func TestExceptsSecondNotSliceShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			return
		} else {
			t.Fail()
		}
	}()
	arrutil.Excepts([]string{"a", "b"}, [1]string{"a"}, arrutil.StringEqualsComparer)
}

func TestExceptsFirstEmptyShouldReturnsEmpty(t *testing.T) {
	data := []string{}
	result := arrutil.Excepts(data, []string{"a", "b"}, arrutil.StringEqualsComparer).([]string)
	assert.Eq(t, []string{}, result)
	assert.NotSame(t, &data, &result, "should always returns new slice")
}

func TestExceptsSecondEmptyShouldReturnsFirst(t *testing.T) {
	data := []string{"a", "b"}
	result := arrutil.Excepts(data, []string{}, arrutil.StringEqualsComparer).([]string)
	assert.Eq(t, data, result)
	assert.NotSame(t, &data, &result, "should always returns new slice")
}

// Intersects tests
func TestIntersectsShouldPassed(t *testing.T) {
	data := []string{
		"a",
		"b",
		"c",
	}
	result := arrutil.Intersects(data, []string{"a", "b"}, arrutil.StringEqualsComparer)
	assert.Eq(t, []string{"a", "b"}, result.([]string))
}

func TestIntersectsFirstNotSliceShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			return
		} else {
			t.Fail()
		}
	}()
	arrutil.Intersects([1]string{"a"}, []string{"a", "b"}, arrutil.StringEqualsComparer)
}

func TestIntersectsSecondNotSliceShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			return
		} else {
			t.Fail()
		}
	}()
	arrutil.Intersects([]string{"a", "b"}, [1]string{"a"}, arrutil.StringEqualsComparer)
}

func TestIntersectsFirstEmptyShouldReturnsEmpty(t *testing.T) {
	data := []string{}
	second := []string{"a", "b"}
	result := arrutil.Intersects(data, second, arrutil.StringEqualsComparer).([]string)
	assert.Eq(t, []string{}, result)
	assert.NotSame(t, &second, &result, "should always returns new slice")
}

func TestIntersectsSecondEmptyShouldReturnsEmpty(t *testing.T) {
	data := []string{"a", "b"}
	second := []string{}
	result := arrutil.Intersects(data, second, arrutil.StringEqualsComparer).([]string)
	assert.Eq(t, []string{}, result)
	assert.NotSame(t, &data, &result, "should always returns new slice")
}

// Union tests

func TestUnionShouldPassed(t *testing.T) {
	data := []string{
		"a",
		"b",
		"c",
	}
	result := arrutil.Union(data, []string{"a", "b", "d"}, arrutil.StringEqualsComparer).([]string)
	assert.Eq(t, []string{"a", "b", "c", "d"}, result)
}

func TestUnionFirstNotSliceShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			return
		} else {
			t.Fail()
		}
	}()
	arrutil.Union([1]string{"a"}, []string{"a", "b"}, arrutil.StringEqualsComparer)
}

func TestUnionSecondNotSliceShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			return
		} else {
			t.Fail()
		}
	}()

	arrutil.Union([]string{"a", "b"}, [1]string{"a"}, arrutil.StringEqualsComparer)
}

func TestUnionFirstEmptyShouldReturnsSecond(t *testing.T) {
	data := []string{}
	second := []string{"a", "b"}
	result := arrutil.Union(data, second, arrutil.StringEqualsComparer).([]string)
	assert.Eq(t, []string{"a", "b"}, result)
	assert.NotSame(t, &second, &result, "should always returns new slice")
}

func TestUnionSecondEmptyShouldReturnsFirst(t *testing.T) {
	data := []string{"a", "b"}
	second := []string{}
	result := arrutil.Union(data, second, arrutil.StringEqualsComparer).([]string)
	assert.Eq(t, data, result)
	assert.NotSame(t, &data, &result, "should always returns new slice")
}

// Find tests
func TestFindShouldPassed(t *testing.T) {
	data := []string{
		"a",
		"b",
		"c",
	}

	result, err := arrutil.Find(data, func(a any) bool { return a == "b" })
	assert.Nil(t, err)
	assert.Eq(t, "b", result)

	_, err = arrutil.Find(data, func(a any) bool { return a == "d" })
	assert.NotNil(t, err)
	assert.Eq(t, arrutil.ErrElementNotFound, err.Error())

}

func TestFindNotSliceShouldPanic(t *testing.T) {
	assert.Panics(t, func() {
		_, _ = arrutil.Find([1]string{"a"}, func(a any) bool { return a == "b" })
	})
}

func TestFindEmptyReturnsErrElementNotFound(t *testing.T) {
	data := []string{}
	_, err := arrutil.Find(data, func(a any) bool { return a == "b" })
	assert.NotNil(t, err)
	assert.Eq(t, arrutil.ErrElementNotFound, err.Error())
}

// FindOrDefault tests
func TestFindOrDefaultShouldPassed(t *testing.T) {
	data := []string{
		"a",
		"b",
		"c",
	}

	result := arrutil.FindOrDefault(data, func(a any) bool { return a == "b" }, "d").(string)
	assert.Eq(t, "b", result)

	result = arrutil.FindOrDefault(data, func(a any) bool { return a == "d" }, "d").(string)
	assert.Eq(t, "d", result)
}

// TakeWhile tests
func TestTakeWhileShouldPassed(t *testing.T) {
	data := []string{
		"a",
		"b",
		"c",
	}

	result := arrutil.TakeWhile(data, func(a any) bool { return a == "b" || a == "c" }).([]string)
	assert.Eq(t, []string{"b", "c"}, result)
}

func TestTakeWhileNotSliceShouldPanic(t *testing.T) {
	assert.Panics(t, func() {
		arrutil.TakeWhile([1]string{"a"}, func(a any) bool { return a == "b" || a == "c" })
	})
}

func TestTakeWhileEmptyReturnsEmpty(t *testing.T) {
	var data []string
	result := arrutil.TakeWhile(data, func(a any) bool { return a == "b" || a == "c" }).([]string)
	assert.Eq(t, []string{}, result)
	assert.NotSame(t, &data, &result, "should always returns new slice")
}

// ExceptWhile tests

func TestExceptWhileShouldPassed(t *testing.T) {
	data := []string{
		"a",
		"b",
		"c",
	}

	result := arrutil.ExceptWhile(data, func(a any) bool { return a == "b" || a == "c" }).([]string)
	assert.Eq(t, []string{"a"}, result)
}

func TestExceptWhileNotSliceShouldPanic(t *testing.T) {
	assert.Panics(t, func() {
		arrutil.ExceptWhile([1]string{"a"}, func(a any) bool { return a == "b" || a == "c" })
	})
}

func TestExceptWhileEmptyReturnsEmpty(t *testing.T) {
	var data []string
	result := arrutil.ExceptWhile(data, func(a any) bool { return a == "b" || a == "c" }).([]string)

	assert.Eq(t, []string{}, result)
	assert.NotSame(t, &data, &result, "should always returns new slice")
}

func TestMap(t *testing.T) {
	list1 := []map[string]any{
		{"name": "tom", "age": 23},
		{"name": "john", "age": 34},
	}

	flatArr := arrutil.Column(list1, func(obj map[string]any) (val any, find bool) {
		return obj["age"], true
	})

	assert.NotEmpty(t, flatArr)
	assert.Contains(t, flatArr, 23)
	assert.Len(t, flatArr, 2)
	assert.Eq(t, 34, flatArr[1])
}
