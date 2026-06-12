package benchmarks

import (
	"strings"
	"testing"
)

func match(subject, target string) bool {
	if !strings.Contains(target, "*") {
		return subject == target
	}

	var sub, tar string
	for subject != "" || target != "" {
		sub, subject, _ = strings.Cut(subject, ".")
		tar, target, _ = strings.Cut(target, ".")
		if sub != tar && tar != "*" {
			return false
		}
	}
	return subject == target
}

func Benchmark_match(b *testing.B) {
	for b.Loop() {
		match("apiservice.match.1.od:match:12345", "apiservice.match.1.*")
	}
	b.ReportAllocs()
}
