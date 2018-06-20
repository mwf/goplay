package bar

import (
	"github.com/mwf/goplay/vgo/vtest/foo"
)

const (
	Bar    = "I'm bar/Bar"
	FooBar = Bar + "; " + foo.Foo
)
