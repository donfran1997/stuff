package main
/*
#cgo CFLAGS: -fplugin=./trust.so

void helloWorld() {
    printf("Hello World\n");
}
*/

import "C"
import "unsafe"

func main() {
	C.helloWorld()
	return
}
