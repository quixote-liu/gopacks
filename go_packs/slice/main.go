package main

import "fmt"

func main() {
	var ss = make([]int, 0)
	appendSlice(&ss)
	appendSliceNoPo(ss)
	fmt.Printf("ss: %v\n", ss)
}

func appendSlice(ss *[]int) {
	*ss = append(*ss, 10)
}

func appendSliceNoPo(ss []int) {
	ss = append(ss, 12)
}
