package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println(arr[2:6])
	fmt.Println(arr[:6])
	s1 := arr[2:]
	updateSlice(s1)
	fmt.Println(s1)
	s2 := arr[:]
	fmt.Println(s2)

	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]

	fmt.Println(s2)

	fmt.Println("Extending slice")
	s1 = arr[2:6]
	s2 = s1[3:5] //s1[3], s1[4]
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)

	//s4 and s5 no longer view arr.
	fmt.Println(arr)

}
