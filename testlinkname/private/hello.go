package private

import _ "unsafe"

//go:linkname hello github.com/ghjan/goskill/testlinkname/hello.hellofunc
func hello() string {
	return "private.hello()"
}
