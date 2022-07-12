package main

import (
	"fmt"
	"gin-demo/scrapy"
)

func main() {
	root := scrapy.Spider{}.BaiduB2b()
	fmt.Println(root)
}
