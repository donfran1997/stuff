package main
/*
#cgo CFLAGS: -fplugin=./trust.so

void helloWorld() {
    printf("Hello World\n");
}
*/

import "C"

func main() {
	C.helloWorld()
	return
}
