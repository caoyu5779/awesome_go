package main

import (
	"fmt"
	"os"
	"bufio"
	"selfLearning/functional/Fib"
	"errors"
)

/*
* 	Open/Close
	Lock/Unlock
	PrintHeader/PrintFooter
*/
//func tryDefer(){
//	//defer fmt.Println(1)
//	//defer fmt.Println(2) //FILO 先进后出
//	//fmt.Println(3)
//	//panic("error")
//	//fmt.Println(4)
//
//	for i :=0; i < 100; i++{
//		defer fmt.Println(i)
//		if i == 30{
//			panic("printed to many")
//		}
//	}
//}

func writeFile(filename string){
	file, err := os.OpenFile(
		filename, os.O_EXCL | os.O_CREATE, 0666)

	err = errors.New("this is a custom error")
	if err != nil {
		//panic(err)
		if pathError, ok := err.(*os.PathError); !ok{
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		//fmt.Println("Error :", err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	defer writer.Flush() //资源管理 close的时候，将操作完成，并结束，在defer 是计算
	f := Fib.Fibonacci()
	for i:= 0; i < 20; i++{
		fmt.Fprintln(writer, f())
	}

}

func main() {
	writeFile("fib.txt")
	//tryDefer()
}
