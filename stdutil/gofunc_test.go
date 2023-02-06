package stdutil_test

import (
	"testing"

	"github.com/zhangyiming748/pretty/stdutil"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

func TestFuncName(t *testing.T) {
	name := stdutil.FuncName(stdutil.PkgName)
	assert.Eq(t, "github.com/zhangyiming748/pretty/stdutil.PkgName", name)

	name = stdutil.FuncName(stdutil.PanicIfErr)
	assert.Eq(t, "github.com/zhangyiming748/pretty/stdutil.PanicIfErr", name)
}

func TestPkgName(t *testing.T) {
	name := stdutil.PkgName(stdutil.FuncName(stdutil.PanicIfErr))
	assert.Eq(t, "github.com/zhangyiming748/pretty/stdutil", name)
}

func TestFullFcName_Parse(t *testing.T) {
	fullName := stdutil.FuncName(stdutil.PanicIfErr)

	ffn := stdutil.FullFcName{FullName: fullName}
	ffn.Parse()
	assert.Eq(t, fullName, ffn.String())
	assert.Eq(t, "stdutil", ffn.PkgName())
	assert.Eq(t, "PanicIfErr", ffn.FuncName())
	assert.Eq(t, "github.com/zhangyiming748/pretty/stdutil", ffn.PkgPath())
	pretty.P(ffn)

	st := stdutil.FullFcName{}
	fullName = stdutil.FuncName(st.FuncName)

	ffn = stdutil.FullFcName{FullName: fullName}
	ffn.Parse()
	assert.Eq(t, "(*FullFcName).FuncName-fm", ffn.FuncName())
	pretty.P(ffn)
}

func TestCutFuncName(t *testing.T) {
	fullName := stdutil.FuncName(stdutil.PanicIfErr)

	pkgPath, funcName := stdutil.CutFuncName(fullName)
	assert.Eq(t, "PanicIfErr", funcName)
	assert.Eq(t, "github.com/zhangyiming748/pretty/stdutil", pkgPath)
}
