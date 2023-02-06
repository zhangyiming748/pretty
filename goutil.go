// Package goutil 💪 Useful utils for Go: int, string, array/slice, map, error, time, format, CLI, ENV, filesystem,
// system, testing, debug and more.
package pretty

import (
	"fmt"

	"github.com/zhangyiming748/pretty/stdutil"
)

// Value alias of stdutil.Value
type Value = stdutil.Value

// Panicf format panic message use fmt.Sprintf
func Panicf(format string, v ...any) {
	panic(fmt.Sprintf(format, v...))
}

// PanicIfErr if error is not empty, will panic
func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

// PanicErr if error is not empty, will panic
func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

// MustOK if error is not empty, will panic
func MustOK(err error) {
	if err != nil {
		panic(err)
	}
}

// PkgName get current package name. alias of stdutil.PkgName()
//
// Usage:
//
//	funcName := goutil.FuncName(fn)
//	pgkName := goutil.PkgName(funcName)
func PkgName(funcName string) string {
	return stdutil.PkgName(funcName)
}

// ErrOnFail return input error on cond is false, otherwise return nil
func ErrOnFail(cond bool, err error) error {
	return OrError(cond, err)
}

// OrError return input error on cond is false, otherwise return nil
func OrError(cond bool, err error) error {
	if !cond {
		return err
	}
	return nil
}

// OrValue get
func OrValue[T any](cond bool, okVal, elVal T) T {
	if cond {
		return okVal
	}
	return elVal
}

// OrReturn call okFunc() on condition is true, else call elseFn()
func OrReturn[T any](cond bool, okFn, elseFn func() T) T {
	if cond {
		return okFn()
	}
	return elseFn()
}
