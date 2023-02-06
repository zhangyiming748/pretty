package assert_test

import (
	"errors"
	"testing"

	"github.com/zhangyiming748/pretty/testutil/assert"
)

func TestAssertions_Chain(t *testing.T) {
	// err := "error message"
	err := errors.New("error message")

	as := assert.New(t).
		NotNil(err).
		Err(err).
		ErrMsg(err, "error message")

	assert.True(t, as.IsOk())
	assert.False(t, as.IsFail())
}
