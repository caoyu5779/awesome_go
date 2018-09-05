package main

import (
	"fmt"
	"regexp"
)

const text = "My email is caoyu5779@126.com\n" +
	"email1 is abc@def.otf\n" +
	"email2 is      blank@qq.com.cn"

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
