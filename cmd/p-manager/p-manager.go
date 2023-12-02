package main

import (
	"fmt"

	diskWalker "github.com/x24870/p-manager/pkg/disk-walker"
)

func main() {
	fmt.Println("P_MANAGER")
	dw := diskWalker.NewDiskWalkerImpl()
	dw.Walk()
}
