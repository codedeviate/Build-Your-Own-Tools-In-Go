package main

import (
	"fmt"
	"os"

	"github.com/username/subsystem/lib"
)

func main() {
	var url string

	if len(os.Args) > 1 {
		url = os.Args[1]
	} else {
		fmt.Print("Enter the URL (e.g., example.com): ")
		fmt.Scanln(&url)
	}

	lib.GetServerInfo(url)
}
