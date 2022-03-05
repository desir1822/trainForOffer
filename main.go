package main

import "fmt"

func main() {
	s := make([]int, 3, 10)
	s = append(s[:], 1)
	fmt.Println(s)
	fmt.Println(len(s))
	t := []byte{}
	fmt.Println(len(t))

}
