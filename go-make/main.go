package main

import (
	"fmt"
	"github.com/username/go-make/vars"
	"os"

	"github.com/username/go-make/cmd"
)

func main() {
	vars.Init()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
