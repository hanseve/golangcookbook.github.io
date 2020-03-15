//Concatenating string benchmark test
//go test -bench=. -benchmem
package main

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkConcat(b *testing.B) {
	var str string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str += "x"
	}
	b.StopTimer()
	if s := strings.Repeat("x", b.N); str != s {
		b.Errorf("Unexpected result: got:%s want:%s", str, s)
	}
}

func BenchmarkBuffer(b *testing.B) {
	var buf bytes.Buffer
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.WriteString("x")
	}
	b.StopTimer()
	if s := strings.Repeat("x", b.N); buf.String() != s {
		b.Errorf("Unexpected result: got:%s want:%s", buf.String(), s)
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	var strBuilder strings.Builder
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strBuilder.WriteString("x")
	}
	b.StopTimer()
	if s := strings.Repeat("x", b.N); strBuilder.String() != s {
		b.Errorf("Unexpected result: got:%s want:%s", strBuilder.String(), s)
	}
}

func BenchmarkCopy(b *testing.B) {
	bs := make([]byte, b.N)
	bl := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bl += copy(bs[bl:], "x")
	}
	b.StopTimer()
	if s := strings.Repeat("x", b.N); string(bs) != s {
		b.Errorf("Unexpected result: got:%s want:%s", string(bs), s)
	}
}

func BenchmarkStringJoin(b *testing.B) {
	var strSlice []string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strSlice = append(strSlice, "x")
	}
	b.StopTimer()
	if s, str := strings.Repeat("x", b.N), strings.Join(strSlice, ""); str != s {
		b.Errorf("Unexpected result: got:%s want:%s", str, s)
	}
}
