---
title: Go 1.14 and Beyond
subtitle: What we know already
author: Andreas Linz—klingt.net
date: 2020-02-21T19+16:11
---

# Go 1.14 {data-background-image=https://images.unsplash.com/photo-1467219598992-52591d77fdec?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1350&q=80 data-background-opacity=0.31415}

--- 

- [preliminary release notes][go114]
- module support in the go command is now **ready for production use**
- they even added [Subversion support][subversion support] for modules 🤣

---

### Embedding of overlapping interfaces 

- [will not compile today](https://play.golang.org/p/SIlnehVqmH-)
- [allowed with Go 1.14][proposalOverlappingInterfaces]

```sh
$ docker run -v $PWD:/app -w /app --rm -it golang:1.13 go run overlapping-interfaces.go
./overlapping-interfaces.go:10:2: duplicate method Close
$ docker run -v $PWD:/app -w /app --rm -it golang:1.14-rc go run overlapping-interfaces.go
works
```

---

### GOINSECURE

- instuct `go` command to skip certificate validation or accept HTTP
- useful when developing a module proxy
- `GOINSECURE=localhost go get foo/bar` 

---

### Vendoring

- if `go.mod` specifies `Go 1.14` and a `/vendor` folder present then commands will default to using vendoring
- usage of modules can be enforced with `-mod=mod`

---

### Test log output

- output of `go test -v` (or `t.Log`) is now streamed instead of being presented at the end of all tests

# Runtime Improvements

--- 

### (almost) Zero Overhead `defer`

> This release improves the performance of most uses of defer to incur almost zero overhead compared to calling the deferred function directly. As a result, defer can now be used in performance-critical code without overhead concerns. 

---

### More Efficient Page Allocator

> The page allocator is more efficient and incurs significantly less lock contention at high values of GOMAXPROCS. This is most noticeable as lower latency and higher throughput for large allocations being done in parallel and at a high rate. 

Discord should consider to reevaluate their switch [from Go to Rust][fromGoToRust] 😉.

---

### New package `hash/maphash`

- collision-resistant
- not cryptographically secure
- hash functions map arbitrary `string`/`[]byte` to 64-bit integers
- useful for building custom maps
- _not safe for concurrent use by multiple goroutines_ (can be initialized with the same seed to get equal hashes)

# Minor Library Changes

---

### `crypto/tls`

- SSL3 support removed
- TLS1.3 can not be disabled using `GODEBUG` anymore
- when multiple [certificate chains](https://tip.golang.org/pkg/crypto/tls/#Config.Certificates) are configured) the first compatible with the peer is selected automatically (e.g. for providing ECDSA and RSA certificates)

---

### `ioutil.TempDir`

- `ioutil.TempDir` allows to specify a naming pattern (like `TempFile`)
- useful when cleaning up temporary stuff of an application (e.g. `rm -rf /tmp/my-app-*`)

```sh
$ docker run -v $PWD:/app -w /app --rm -it golang:1.14-rc go run predictable-tempdir.go
/tmp/584949244-my-app
/tmp/my-app-532238635
```

---

### `mime`

- default type of `.js` and `.mjs` files is now `text/javascript`
- `application/javascript` deprecated ([IETF draft](https://datatracker.ietf.org/doc/draft-ietf-dispatch-javascript-mjs/))

### `net/http`

- `Header.Values` returns all values associated to a header key (`Header.Value` only returns the first one)
- `httptest/Server` adds `EnableHTTP2`

### `testing`

- `t.Cleanup` allows to define a custom cleanup function that is run after each test

```sh
$ docker run -v $PWD:/app -w /app --rm -it golang:1.14-rc go test -v ./cleanup-test-artifacts_test.go
=== RUN   TestCleanup
    TestCleanup: cleanup-test-artifacts_test.go:15: removing  /tmp/golang-leipzig-631932423
--- PASS: TestCleanup (0.00s)
PASS
```

# Go 1.15 {data-background-image=https://images.unsplash.com/photo-1519567141891-788b756572ab?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1863&q=80 data-background-opacity=0.31415}

---

- [proposals][go115]
- scheduled for August 2020
- no major changes

---

### Int to String

- `cmd/vet` warn about [`int` to `string` conversion][issue32479]
- just replace those `int`s with `rune`s
- example: `string(10)` is `"\n"` not `"10"`
- `unicode/utf8` provides `runeToString` ([play](https://play.golang.org/p/ZnUF0Oc_dAG))

---

### Impossible `interface`-`interface` assertion

> Currently, Go permits any type assertion x.(T) (and corresponding type switch case) where the type of x and T are interfaces.

---

- example: Two methods with the same name buth different signature ([play](https://play.golang.org/p/ah551xs4So0))
- impossible conversion is [known at compile-time][issue4483]
- `cmd/vet` will warn about this (will be a compile error later)

### Constant-evaluate index and slice expressions

- if the indices and values are constant then this [can be done at compile time][issue28591]
- turns out it is more complicated and would be [backwards incompatible][issue28591comment]

# Future  {data-background-image=https://images.unsplash.com/photo-1543083115-638c32cd3d58?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=2689&q=80 data-background-opacity=0.31415}

---

> The primary goals for Go remain package and version management, better error handling support, and generics. 

- [try proposal](https://golang.org/issue/32437) was abandoned because of controversy

```go
f, err := os.Open(filename)
if err != nil {
	return …, err  // zero values for other results, if any
}
// with try
f := try(os.Open(filename))
```

> we are also making progress on the generics front (more on that later this year).

[go114]: https://tip.golang.org/doc/go1.14
[go115]: https://blog.golang.org/go1.15-proposals
[issue32479]: https://github.com/golang/go/issues/32479
[issue4483]: https://golang.org/issue/4483
[issue28591]: https://github.com/golang/go/issues/28591
[issue28591comment]: https://github.com/golang/go/issues/28591#issuecomment-579993684
[proposalOverlappingInterfaces]: https://github.com/golang/proposal/blob/master/design/6977-overlapping-interfaces.md
[subversion support]: https://go-review.googlesource.com/c/go/+/203497/
[fromGoToRust]: https://blog.discordapp.com/why-discord-is-switching-from-go-to-rust-a190bbca2b1f