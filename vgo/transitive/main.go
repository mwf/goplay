package main

import (
	"github.com/mwf/goplay/vgo/vtest/bar"
	"github.com/mwf/goplay/vgo/vtest/foo"
)

func main() {
	println("bar", bar.Bar)
	println("foobar", bar.FooBar)
	println("foo", foo.Foo)
}
