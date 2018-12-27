package main

import (
	"fmt"

	"github.com/zjutdp/mathy" // 引用刚刚创建的包名
)

func main() {
	fmt.Println(mathy.Fib(10))
	fmt.Println(mathy.Fact(10))
}
