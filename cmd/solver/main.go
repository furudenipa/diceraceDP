package main

import (
	"fmt"
	"os"

	"github.com/furudenipa/diceraceDP/pkg/mdpsolver"
)

func main() {
	policyFilePath := "./data/policy_.bin"
	if len(os.Args) > 1 {
		policyFilePath = os.Args[1]
	}

	fmt.Printf("Save to: %s\n", policyFilePath)
	mdpsolver.Run(policyFilePath)
}
