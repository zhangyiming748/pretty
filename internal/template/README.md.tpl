# Go Util

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/zhangyiming748/pretty?style=flat-square)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/zhangyiming748/pretty)](https://github.com/zhangyiming748/pretty)
[![Go Report Card](https://goreportcard.com/badge/github.com/zhangyiming748/pretty)](https://goreportcard.com/report/github.com/zhangyiming748/pretty)
[![Unit-Tests](https://github.com/zhangyiming748/pretty/workflows/Unit-Tests/badge.svg)](https://github.com/zhangyiming748/pretty/actions)
[![Coverage Status](https://coveralls.io/repos/github/zhangyiming748/pretty/badge.svg?branch=master)](https://coveralls.io/github/zhangyiming748/pretty?branch=master)
[![Go Reference](https://pkg.go.dev/badge/github.com/zhangyiming748/pretty.svg)](https://pkg.go.dev/github.com/zhangyiming748/pretty)

💪 Useful utils package for the Go: int, string, array/slice, map, error, time, format, CLI, ENV, filesystem, system, testing and more.

> **[中文说明](README.zh-CN.md)**

**Basic packages:**

- [`arrutil`](./arrutil): Array/Slice util functions. eg: check, convert, formatting, enum, collections
- [`cliutil`](./cliutil) Command-line util functions. eg: colored print, read input, exec command
- [`envutil`](./envutil) ENV util for current runtime env information. eg: get one, get info, parse var
- [`fmtutil`](./fmtutil) Format data util functions. eg: data, size, time
- [`fsutil`](./fsutil) Filesystem util functions, quick create, read and write file. eg: file and dir check, operate
- [`jsonutil`](./jsonutil) some util functions for quick read, write, encode, decode JSON data.
- [`maputil`](./maputil) Map data util functions. eg: convert, sub-value get, simple merge
- [`mathutil`](./mathutil) Math(int, number) util functions. eg: convert, math calc, random
- `netutil` Network util functions
- [`reflects`](./reflects) Provide extends reflect util functions.
- [`stdutil`](./stdutil) Provide some commonly std util functions.
- [`structs`](./structs) Provide some extends util functions for struct. eg: tag parse, struct data init
- [`strutil`](./strutil) String util functions. eg: bytes, check, convert, encode, format and more
- [`sysutil`](./sysutil) System util functions. eg: sysenv, exec, user, process
  - [process](./sysutil/process) Provide some process handle util functions.

**Advance packages:**

- [`cflag`](./cflag):  Wraps and extends go `flag.FlagSet` to build simple command line applications
- cli util:
  - [cmdline](./cliutil/cmdline) Provide cmdline parse, args build to cmdline
- [`dump`](./dump): GO value printing tool. print slice, map will auto wrap each element and display the call location
- [`errorx`](./errorx) Provide an enhanced error implements for go, allow with stacktrace and wrap another error.
- net util:
  - [httpreq](netutil/httpreq) An easier-to-use HTTP client that wraps http.Client
- string util:
  - [textscan](strutil/textscan) Implemented a parser that quickly scans and analyzes text content. It can be used to parse INI, Properties and other formats
- sys util:
  - [clipboard](./sysutil/clipboard) Provide a simple clipboard read and write operations.
  - [cmdr](./sysutil/cmdr) Provide for quick build and run a cmd, batch run multi cmd tasks
- [`testutil`](./testutil) Test help util functions. eg: http test, mock ENV value
  - [assert](./testutil/assert) Asserts functions for help testing
- [`timex`](./timex) Provides an enhanced time.Time implementation. Add more commonly used functional methods
  - such as: DayStart(), DayAfter(), DayAgo(), DateFormat() and more.

## Go Doc

Please see [Go doc](https://pkg.go.dev/github.com/zhangyiming748/pretty)

## Install

```shell
go get github.com/zhangyiming748/pretty
```

## Usage

```go
// github.com/zhangyiming748/pretty
is.True(goutil.IsEmpty(nil))
is.False(goutil.IsEmpty("abc"))

is.True(goutil.IsEqual("a", "a"))
is.True(goutil.IsEqual([]string{"a"}, []string{"a"}))
is.True(goutil.IsEqual(23, 23))

is.True(goutil.Contains("abc", "a"))
is.True(goutil.Contains([]string{"abc", "def"}, "abc"))
is.True(goutil.Contains(map[int]string{2: "abc", 4: "def"}, 4))

// convert type
str = goutil.String(23) // "23"
iVal = goutil.Int("-2") // 2
i64Val = goutil.Int64("-2") // -2
u64Val = goutil.Uint("2") // 2
```

## Packages
{{pgkFuncs}}
## Code Check & Testing

```bash
gofmt -w -l ./
golint ./...

# testing
go test -v ./...
go test -v -run ^TestErr$
go test -v -run ^TestErr$ ./testutil/assert/...
```

Testing in docker:

```shell
cd goutil
docker run -ti -v $(pwd):/go/work golang:1.18
root@xx:/go/work# go test ./...
```

## Related

- https://github.com/duke-git/lancet
- https://github.com/samber/lo
- https://github.com/zyedidia/generic
- https://github.com/thoas/go-funk

## zhangyiming748/pretty packages

- [zhangyiming748/pretty/ini](https://github.com/zhangyiming748/pretty/ini) Go config management, use INI files
- [zhangyiming748/pretty/rux](https://github.com/zhangyiming748/pretty/rux) Simple and fast request router for golang HTTP
- [zhangyiming748/pretty/gcli](https://github.com/zhangyiming748/pretty/gcli) Build CLI application, tool library, running CLI commands
- [zhangyiming748/pretty/slog](https://github.com/zhangyiming748/pretty/slog) Lightweight, easy to extend, configurable logging library written in Go
- [zhangyiming748/pretty/color](https://github.com/zhangyiming748/pretty/color) A command-line color library with true color support, universal API methods and Windows support
- [zhangyiming748/pretty/event](https://github.com/zhangyiming748/pretty/event) Lightweight event manager and dispatcher implements by Go
- [zhangyiming748/pretty/cache](https://github.com/zhangyiming748/pretty/cache) Generic cache use and cache manager for golang. support File, Memory, Redis, Memcached.
- [zhangyiming748/pretty/config](https://github.com/zhangyiming748/pretty/config) Go config management. support JSON, YAML, TOML, INI, HCL, ENV and Flags
- [zhangyiming748/pretty/filter](https://github.com/zhangyiming748/pretty/filter) Provide filtering, sanitizing, and conversion of golang data
- [zhangyiming748/pretty/validate](https://github.com/zhangyiming748/pretty/validate) Use for data validation and filtering. support Map, Struct, Form data
- [zhangyiming748/pretty](https://github.com/zhangyiming748/pretty) Some utils for the Go: string, array/slice, map, format, cli, env, filesystem, test and more
- More, please see https://github.com/zhangyiming748/pretty

## License

[MIT](LICENSE)
