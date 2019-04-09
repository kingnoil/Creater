package main

import (
	"fmt"
	"regexp"
)

const text  =  `
My email is yangzi@zuoyebang.com@abc.com
email1 is abc@1.com
email2 is kkk@qq.com
email3 is ddd@abc.com.cn
`
func main() {
	re := regexp.MustCompile(
		`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)(\.[a-zA-Z0-9]+)`)
	match  := re.FindAllStringSubmatch(text,-1)
	for v,m := range match{
		fmt.Println(v,m)
	}

}
