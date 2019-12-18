package main
//
//#cgo CFLAGS: -fplugin=./trust.so
//
//void test() {
//    printf("Hello World");
//}
//
//
import "C"

func main() {
	C.test()
	return
}
