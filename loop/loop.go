package main

import (
	"io"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	printFile("abc.txt")

	s := `abc"d"
	kkk
	124
	
	p`

	printFileContents(strings.NewReader(s))
}

func printFile(filename string){
	file,err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}