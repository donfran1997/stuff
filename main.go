package main

/*
#cgo CFLAGS: -fplugin=./trust.so

void hello() {
    printf("Hello World");
}
*/

import "C"
import "unsafe"

func main() {
	C.hello()
	return
}
