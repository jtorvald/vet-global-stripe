//go:build !go1.12
// +build !go1.12

package main

import (
	"github.com/jtorvald/vet-global-stripe/passes/globalstripe"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(globalstripe.Analyzer) }
