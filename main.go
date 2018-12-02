package main

import (
	"github.com/ScottHuangZL/blockchain/cli"
	"os"
)


func main() {
	// fmt.Println("Hello Blockchain from Scott Huang")
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()

}
