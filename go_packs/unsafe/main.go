package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// slice header
	fmt.Println("===============sliceHeader==============")
	sliceHeader()

	// hmap()
	fmt.Println("==============hamp===================")
	hmap()

	fmt.Println("==============OffsetOf=================")
	OffsetOf()

	fmt.Println("============StringToMap===============")
	bytes := stringToBytes("liuchengshun")
	fmt.Println(bytes)                                 // [108 105 117 99 104 101 110 103 115 104 117 110]
	fmt.Println(string(stringToBytes("liuchengshun"))) // liuchengshun
	fmt.Println(bytesToString(bytes))                  // liuchengshun
}

// the define of slice structure:
// runtime/slice.go
// type slice struct {
//     array unsafe.Pointer // 元素指针
//     len   int // 长度
//     cap   int // 容量
// }

func sliceHeader() {
	s := make([]int, 9, 20)
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println(Len, len(s)) // 9, 9

	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(Cap, cap(s)) // 20, 20
}

// the define of map structure:
// type hmap struct {
// 	count     int
// 	flags     uint8
// 	B         uint8
// 	noverflow uint16
// 	hash0     uint32

// 	buckets    unsafe.Pointer
// 	oldbuckets unsafe.Pointer
// 	nevacuate  uintptr

// 	extra *mapextra
// }

func hmap() {
	m := make(map[string]int)
	m["name"] = 957
	m["id"] = 000

	count := **(**int)(unsafe.Pointer(&m))
	fmt.Println(count, len(m)) // 2, 2
}

type Programmer struct {
	name     string
	language string
}

func OffsetOf() {
	p := Programmer{"lcs", "php"}
	fmt.Println(p)

	name := (*string)(unsafe.Pointer(&p))
	*name = "lcs2"

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
	*lang = "golang"

	fmt.Println(p)
}

func stringToBytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	sliceHeader := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&sliceHeader))
}

func bytesToString(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	stringHeader := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}

	return *(*string)(unsafe.Pointer(&stringHeader))
}
