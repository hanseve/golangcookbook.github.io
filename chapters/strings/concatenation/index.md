---
title: Concatenating Strings
question: How do I combine strings?
---
There are different to Concatenating strings:

1.Use operator `+`

{% include example.html example="concat" %}

2.Use `bytes.Buffer`

{% include example.html example="buffer" %}

3.Use `strings.Builder`, the efficient way, since Go 1.10 

{% include example.html example="string_builder" %}

4.Use `copy()`

{% include example.html example="copy" %}

5.Use `append(), strings.Join()`

{% include example.html example="string_join" %}

Run test
```
$ go test -bench=. -benchmem Condatenation_test.go

goos: darwin
goarch: amd64
BenchmarkConcat-8                1000000             39870 ns/op          503994 B/op          1 allocs/op
BenchmarkBuffer-8               200000000                8.86 ns/op            2 B/op          0 allocs/op
BenchmarkStringBuilder-8        1000000000               4.58 ns/op            6 B/op          0 allocs/op
BenchmarkCopy-8                 300000000                4.44 ns/op            0 B/op          0 allocs/op
BenchmarkStringJoin-8           20000000                89.2 ns/op            80 B/op          0 allocs/op
PASS

```

> Reference:
> [https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go](https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go)



