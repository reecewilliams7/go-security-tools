package main

import (
	"fmt"

	cmds "github.com/reecewilliams7/go-security-tools/cmd/gst/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	err := doc.GenMarkdownTree(cmds.RootCmd, "./docs")
	if err != nil {
		fmt.Println(err.Error())
	}
}
