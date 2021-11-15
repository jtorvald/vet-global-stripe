//go:build go1.12
// +build go1.12

package main

import (
	"github.com/jtorvald/vet-global-stripe/passes/globalstripe"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/unitchecker"
)

// Analyzers returns analyzers of bodyclose.
func analyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		globalstripe.Analyzer,
	}
}

func main() {
	unitchecker.Main(analyzers()...)
}
