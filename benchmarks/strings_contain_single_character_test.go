package benchmarks

import (
	"strings"
	"testing"
)

type input struct {
	text  string
	found bool
}

var examples = []input{
	{text: "topic", found: false},
	{text: "*", found: true},
	{text: "topic.match.1", found: false},
	{text: "topic.*.1", found: true},
	{text: "topic.match.2", found: false},
	{text: "topic.match.*", found: true},
	{text: "topic.match.1.update.5.suspend", found: false},
	{text: "topic.match.1.update.*.suspend", found: true},
	{text: "topic.match.1.update.6.open", found: false},
	{text: "topic.match.1.*.*.*", found: true},
	{text: "topic.match.*.update.*.suspend", found: true},
	{text: "topic.match.*.update.*.*", found: true},
}

func Benchmark_strings_Contains(b *testing.B) {
	for b.Loop() {
		for _, example := range examples {
			_ = strings.Contains(example.text, "*")
		}
	}
}

func Benchmark_strings_ContainsRune(b *testing.B) {
	for b.Loop() {
		for _, example := range examples {
			_ = strings.ContainsRune(example.text, '*')
		}
	}
}

func Benchmark_strings_ContainsFunc_λInLoop(b *testing.B) {
	for b.Loop() {
		for _, example := range examples {
			_ = strings.ContainsFunc(example.text, func(r rune) bool { return r == '*' })
		}
	}
}

func Benchmark_strings_ContainsFunc_preparedλ(b *testing.B) {
	fn := func(r rune) bool { return r == '*' }

	for b.Loop() {
		for _, example := range examples {
			_ = strings.ContainsFunc(example.text, fn)
		}
	}
}

/*
$ go test -bench=. -benchtime=10s -benchmem ./benchmarks/...
goos: linux
goarch: amd64
pkg: github.com/Radek-Pysny/go-tests/benchmarks
cpu: AMD Ryzen 7 PRO 7840U w/ Radeon 780M Graphics
Benchmark_strings_Contains-16                        101005335   105.6 ns/op    0 B/op   0 allocs/op
Benchmark_strings_ContainsRune-16                    216583808    55.38 ns/op   0 B/op   0 allocs/op
Benchmark_strings_ContainsFunc_λInLoop-16            41462899    311.3 ns/op    0 B/op   0 allocs/op
Benchmark_strings_ContainsFunc_preparedλ-16          39162380    305.3 ns/op    0 B/op   0 allocs/op
PASS
ok      github.com/Radek-Pysny/go-tests/benchmarks   47.538s
*/
