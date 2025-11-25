package main

import "fmt"

func main() {
	if a := 1; a > 0 {
		fmt.Println(a)
	}
	c := make(map[string]string)
	c["a"] = "anda"
	val, ok := c["b"]
	fmt.Println(val, ok)
	b := "hello"
	fmt.Println(b[:2])
}
