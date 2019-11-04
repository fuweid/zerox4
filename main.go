package main

import (
	"fmt"

	"github.com/containerd/cgroups"
)

func main() {
	fmt.Println(cgroups.NewCpu(""))
}
