package main

import (
	"fmt"

	"github.com/galaxy-toolkit/ippool/internal/global"
)

func main() {
	fmt.Println(global.Config)

	global.Logger.Info("hello world")
}
