package main

import (
	"fmt"
	"os"
	"github.com/p3tr0v/chacal/antidebug"
	"github.com/p3tr0v/chacal/antimem"
	"github.com/p3tr0v/chacal/antivm"
)

func main() {
	if antidebug.ByProcessWatcher() {
		os.Exit(1)
	} else if antimem.ByMemWatcher() {
		os.Exit(1)
	} else if antivm.BySizeDisk(100) {
		os.Exit(1)
	} else if antivm.IsVirtualDisk() {
		os.Exit(1)
	} else if antivm.ByMacAddress() {
		os.Exit(1)
	} else {
		fmt.Println("Malicious Code Working !")
	}
}
