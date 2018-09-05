package main

import (
	"fmt"
	"io/ioutil"
)

//func eval(a, b int, op string) int {
//	var result int
//	switch op {
//	case "+":
//		result = a + b
//	case "-":
//		result = a - b
//	case "*":
//		result = a * b
//	case "/":
//		result = a / b
//	default:
//		panic("unsupport op : " + op)
//	}
//
//	return result
//}

func main() {
	const filename = "abc.txt"

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	result, err := eval(3, 4, "/")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}
