# What is it?

A test case for `vgo vendor` bug

# How to reproduce

```bash
go get -u golang.org/x/vgo

vgo vendor
vgo: finding github.com/stretchr/testify v1.2.2
vgo: finding github.com/stretchr/objx v0.0.0-20180106011353-facf9a85c22f
vgo: finding github.com/stretchr/testify v1.2.0
vgo: finding github.com/stretchr/objx v0.0.0-20140526180921-cbeaeb16a013
vgo: finding github.com/pmezard/go-difflib v0.0.0-20151028094244-d8ed2627bdf0
vgo: finding github.com/pmezard/go-difflib v1.0.0
vgo: finding github.com/davecgh/go-spew v1.1.0
vgo: downloading github.com/stretchr/testify v1.2.2
vgo: downloading github.com/davecgh/go-spew v1.1.0
vgo: downloading github.com/pmezard/go-difflib v1.0.0
vgo: import "github.com/davecgh/go-spew/spew/testdata" [/Users/ikorolev/dev/go/src/v/github.com/davecgh/go-spew@v1.1.0/spew/testdata]: no Go source files

vgo test all
ok		github.com/davecgh/go-spew/spew 		0.013s
ok		github.com/mwf/goplay/vgo/testify		0.034s
ok		github.com/pmezard/go-difflib/difflib	0.013s
ok		github.com/stretchr/testify/assert		0.047s

```

`vgo` fails to populate vendor folder while all tests pass.
