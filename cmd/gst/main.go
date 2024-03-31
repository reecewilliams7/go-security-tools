package main

import (
	"fmt"

	"github.com/reecewilliams7/go-security-tools/cmd/gst/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}
