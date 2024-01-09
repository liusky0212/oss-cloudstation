package main

import (
	"fmt"
	"oss-station/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
