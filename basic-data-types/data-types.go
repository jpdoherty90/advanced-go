package main

import (
	"fmt"
	"unsafe"
)


// Return string length without use "len"
func strLenSafe(s string) int  {
	count := 0
	for range s {
		count++
	}
	return count
}

func strLenUnsafe(s string) int  {
	return *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
}

// Return y coordinate of a point without using "p.y"
type Point struct { x , y int }

func getY(p Point) int {
	return *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + uintptr(8)))
}


// Return sum of []int without using range or []
func sumIntArray(a []int) int {
	sum := 0
	intSize := unsafe.Sizeof(int(0))
	arrLen := len(a)
	ap := (uintptr((*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)))))))
	for i := 0; i < arrLen; i++ {
		sum += *(*int)(unsafe.Pointer(ap + uintptr(i)*intSize))
	}
	return sum
}



// Given a map[int]int return the max value without using range or []
func getMapMaxVal(m map[int]int) int {

	// *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + uintptr(8)))

	mp := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&m)) + uintptr(16)))
	fmt.Println(mp)

	// Some ideas:
	// You can get the len in the header
	// You have the pointer to the buckets
	// You can use for len to iterate through values in buckets
	// Maybe need a nested for loop?
	// 


	return 1
}

func main() {

	// Try out string length
	// s := "hello, world"
	// l := strLenSafe(s)
	// fmt.Println(l)
	// l = strLenUnsafe(s)
	// fmt.Println(l)

	// Try out getting y coordinate
	// p := Point{5, 3}
	// y := getY(p)
	// fmt.Println(y)
	
	// Try out array sum
	a := []int{2, 3, 5, 7, 11}
	sum := sumIntArray(a)
	fmt.Println(sum)

	// Try out map max value
	m := map[int]int{
		1:	435,
		2:	824,
		3:	234,
	}
	max := getMapMaxVal(m)
	fmt.Println(max)

}