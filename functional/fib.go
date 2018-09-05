package main

import (
	"bufio"
	"fmt"
	"io"
	"selfLearning/functional/Fib"
	"strings"
)

// 1,1,2,3,5,8,13, ...
// a,b
//   a,b

type intGen func() int

func (g intGen) Read(
	p []byte) (n int, err error) {
	next := g()

	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	//f := fibonacci()
	var f intGen = Fib.Fibonacci()

	printFileContents(f)
	//fmt.Println(f()) //1
	//fmt.Println(f()) //1
	//fmt.Println(f()) //2
	//fmt.Println(f()) //3
	//fmt.Println(f()) //5
	//fmt.Println(f()) //8
	//fmt.Println(f()) //13
	//fmt.Println(f()) //21
}
