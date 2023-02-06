package fmtutil_test

import (
	"testing"

	"github.com/zhangyiming748/pretty/fmtutil"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

func TestHowLongAgo(t *testing.T) {
	tests := []struct {
		args int64
		want string
	}{
		{-36, "unknown"},
		{36, "36 secs"},
		{346, "5 mins"},
		{3467, "57 mins"},
		{346778, "4 days"},
		{1200346778, "13892 days"},
	}

	for _, tt := range tests {
		assert.Eq(t, tt.want, fmtutil.HowLongAgo(tt.args))
	}
}
