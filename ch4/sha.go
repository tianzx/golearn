package main

import (
	"fmt"
)

func reverse(s []int) {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {

		s[i], s[j] = s[j], s[i]

	}
}
func main2() {
	// q := []byte("x")
	// c1 := sha256.Sum256([]byte("x"))
	// c2 := sha256.Sum256([]byte("X"))
	// // c3 := []byte{1, 23, 4}
	// fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// fmt.Printf("%v\n%v\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// fmt.Println(q[0])
	// r := [...]int{99: -1}
	// fmt.Println(r)
	r := [...]int{1, 2, 3, 4, 5}
	fmt.Println(r)
	reverse(r[:])
	fmt.Println(r)
}
