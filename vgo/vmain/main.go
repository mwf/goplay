package main

import (
	vtest2 "github.com/peterbourgon/vtest/v2"
	foo2 "github.com/peterbourgon/vtest/v2/foo"
	vtest3 "github.com/peterbourgon/vtest/v3"
	foo3 "github.com/peterbourgon/vtest/v3/foo"
)

func main() {
	println("vtest2", vtest2.Version, "foo2", foo2.Version)
	println("vtest3", vtest3.Version, "foo3", foo3.Version)
}
