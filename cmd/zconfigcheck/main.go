package main

import (
	"github.com/ccoveille/zconfigcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(zconfigcheck.Analyzer)
}
