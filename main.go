package main

import (
	"os"
	"github.com/ScottHuangZL/blockchain/cli"
	//"./cli"
)

func main() {
	// fmt.Println("Hello Blockchain from Scott Huang")
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
