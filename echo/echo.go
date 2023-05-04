package main

import (
	"fmt"
	"os"
)

func echo1() {
	var sep, s string
	for _, sep = range os.Args {
		s += sep + " "
	}
	fmt.Println(s)
}

func echo2() {
	var sep string
	var index int
	for index, sep = range os.Args[1:] {
		fmt.Printf("索引：%d  参数：%s\n", index, sep)
	}
}
func main() {
	echo2()
	// var s, sep string
	// for i := 0; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)
}
