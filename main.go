package main

import (
	"os"
	"study-golang/numcli"
)

func main() {

	// 型変換を実行
	numcli.Convert(os.Args[1], os.Args[2:])

}
