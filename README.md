sembench
========

Simple Go benchmarks for the "Effective Go" semaphore (with channels)
and a condition/mutex pair.

`go version go1.3.3 darwin/amd64`

    ❯ go test -bench . -benchmem -benchtime 10s
    PASS
    BenchmarkChannelSemaphore    100000000           234 ns/op           0 B/op           0 allocs/op
    BenchmarkCondSemaphore    100000000           175 ns/op           0 B/op           0 allocs/op

