package main

import (
	"os"

	"github.com/furudenipa/diceraceDP/pkg/visualizer"
)

func main() {
	policyFilePath := "../../data/policy3.bin"
	if len(os.Args) > 1 {
		policyFilePath = os.Args[1]
	}
	visualizer.Run(policyFilePath)
}
