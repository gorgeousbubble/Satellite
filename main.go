package main

import (
	"fmt"
	"satellite/pack"
)

func main() {
	fmt.Println("hello,world!")

	src := make([]string, 2)
	src[0] = "d:/Alopex/BadApple!!/AKB-040A.wmv"
	src[1] = "d:/Alopex/BadApple!!/AKB-040B.wmv"
	dest := "e:/Packet/akb40.pak"
	err := pack.PackAES(src, dest)
	if err != nil {
		fmt.Println("packet error.")
	}
	fmt.Println("packet success.")
}
