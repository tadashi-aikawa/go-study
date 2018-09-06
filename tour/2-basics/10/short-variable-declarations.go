package main

import "fmt"

// これは関数外なのでNG
// gg := "gg"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
