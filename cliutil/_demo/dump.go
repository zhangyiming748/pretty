package main

import (
	"os"

	"github.com/zhangyiming748/pretty/dump"
)

func main() {
	dump.P(os.Args)
}
