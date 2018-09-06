package main

import (
	"fmt"
	"strings"
)

func main() {
	file := "hoge^huga.csv"
	ret := strings.Replace(strings.Split(file, ".csv")[0], "^", "", -1)

	fmt.Println(ret)
}
