package main

import "github.com/kjdv/kgul/runtime/assert"

func main() {
	// assert.Assert(false)
	// assert.Assert(false, "asserts ", "with", " arguments")
	assert.Assertf(false, "asserts with some message with %d %s", 2, "items")
	// assert.Failf("blaah %f", 3.14)
}
