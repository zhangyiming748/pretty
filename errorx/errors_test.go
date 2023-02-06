package errorx_test

import (
	"fmt"
	"testing"

	"github.com/zhangyiming748/pretty/errorx"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

func TestErrorR_usage(t *testing.T) {
	err := errorx.NewR(405, "param error")

	assert.Eq(t, 405, err.Code())
	assert.Eq(t, "param error", err.Error())
	assert.Eq(t, "param error(code: 405)", err.String())
	assert.False(t, err.IsSuc())
	assert.True(t, err.IsFail())

	fmt.Println(err)
	fmt.Printf("%v\n", err)
	fmt.Printf("%+v\n", err)
	fmt.Printf("%#v\n", err)

	err = errorx.Suc("ok")
	assert.Eq(t, 0, err.Code())
	assert.True(t, err.IsSuc())
	assert.False(t, err.IsFail())

	err = errorx.Fail(1301, "fail")
	assert.Eq(t, 1301, err.Code())
	assert.False(t, err.IsSuc())
	assert.True(t, err.IsFail())
	assert.NotEmpty(t, err.String())
}

func TestErrMap_usage(t *testing.T) {
	em := make(errorx.ErrMap)
	assert.Nil(t, em.ErrorOrNil())
	assert.Nil(t, em.One())
	assert.True(t, em.IsEmpty())

	em["err1"] = errorx.Raw("this is error1")
	assert.False(t, em.IsEmpty())
	assert.NotEmpty(t, em.Error())
}

func TestErrors_usage(t *testing.T) {
	es := make(errorx.Errors, 0)
	assert.Nil(t, es.First())
	assert.Nil(t, es.ErrorOrNil())
	assert.True(t, es.IsEmpty())

	es = append(es, errorx.Raw("this is error1"))
	assert.False(t, es.IsEmpty())
	assert.NotEmpty(t, es.Error())
}
