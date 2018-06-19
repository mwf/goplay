# What is it?

A test project for adding `vgo` support to github.com/go-chi/chi

# How to play with it

```bash
go get -u golang.org/x/vgo

GOCACHE=off vgo test all

ok      github.com/go-chi/chi/v3    2.967s
ok      github.com/go-chi/chi/v3/middleware 23.319s
?       github.com/mwf/goplay/vgo/chi   0.117s [no test files]
ok      golang.org/x/net/http/httpguts  0.035s
ok      golang.org/x/net/http2  74.363s
ok      golang.org/x/net/http2/hpack    0.051s
ok      golang.org/x/net/idna   0.011s
?       golang.org/x/text/internal/gen  0.124s [no test files]
?       golang.org/x/text/internal/testtext 0.171s [no test files]
ok      golang.org/x/text/internal/ucd  0.010s
ok      golang.org/x/text/secure/bidirule   0.134s
ok      golang.org/x/text/transform 0.117s
ok      golang.org/x/text/unicode/bidi  0.105s
ok      golang.org/x/text/unicode/cldr  0.092s
ok      golang.org/x/text/unicode/norm  2.178s
ok      golang.org/x/text/unicode/rangetable    1.769s

vgo run main.go
Running at localhost:3333
^C

vgo mod -vendor
go run main.go
Running at localhost:3333
```

Using `vgo mod -vendor` we can ensure that old Go version (go 10.x for example) work correctly.
