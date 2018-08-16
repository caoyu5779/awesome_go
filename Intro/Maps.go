package main

import "fmt"

func main() {
	m := map[string]string{ // hash map key 无序, 手动对 key 排序
		"james": "bond",
		"alone": "keys",
		"steve": "jobs",
		"red":   "bag",
	}

	m2 := make(map[string]int) //m2 == empty map

	var m3 map[string]int //m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["james"] // 变量 错误后 获取到 zero value
	fmt.Println(courseName, ok)

	if jamesName, ok := m["james"]; ok {
		fmt.Println(jamesName)
	} else {
		fmt.Println("Key doesnot exit")
	}

	fmt.Println("deleting values")
	name, yes := m["alone"]
	delete(m, "keys")
	fmt.Println(name, yes)

}
