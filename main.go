package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	pwdTable = "abcdefhjmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWYXZ"
	symTable = "!$%&'()*+,-./:;<=>?@[#]^_`{|}~"
	numTable = "23456789"
	length   = flag.Int("len", 16, "密码长度")
	sym      = flag.Bool("sym", false, "启用符号")
	num      = flag.Bool("num", true, "启用数字")
)

func main() {
	flag.Parse()

	// 开启符号
	if *sym {
		pwdTable += symTable
	}

	// 开启数字
	if *num {
		pwdTable += numTable
	}

	// todo 可以指定数字和符号的长度
	pwd := ""
	rand.Seed(int64(time.Now().UnixNano()))
	lenTable := len(pwdTable)
	for i := 0; i < *length; i++ {
		r := rand.Intn(lenTable)
		pwd += pwdTable[r : r+1]
	}

	fmt.Println(pwd)
}
