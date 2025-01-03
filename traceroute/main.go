// main.go
package main

import (
	"flag" // Import the flag package to handle command-line arguments
	"fmt"  // Import fmt for formatted I/O
	"os"   // Import os to handle exit codes and file operations

	// Import the traceroute package that contains the actual traceroute logic
	traceroute "github.com/username/traceroute/traceroutemod"
)

func main() {
	// Define and parse command-line flags:
	// - `-d` for DNS lookup (whether to resolve hostnames)
	// - `-h` for maximum number of hops (default is 30)
	// - `-t` for timeout in seconds (default is 3 seconds)
	config := traceroute.Config{
		Dns:      flag.Bool("d", false, "Do DNS lookup"),     // Boolean flag for DNS lookup
		Hops:     flag.Int("h", 30, "Max hops"),              // Integer flag for max hops
		Timeout:  flag.Float64("t", 3, "Timeout in seconds"), // Float flag for timeout duration in seconds
		SkipHop1: flag.Bool("s1", false, "Skip hop 1"),       // Boolean flag to skip hop 1
		SkipHop2: flag.Bool("s2", false, "Skip hop 2"),       // Boolean flag to skip hop 2
		SkipHop3: flag.Bool("s3", false, "Skip hop 3"),       // Boolean flag to skip hop 3
		SkipHop4: flag.Bool("s4", false, "Skip hop 4"),       // Boolean flag to skip hop 4
		SkipHop5: flag.Bool("s5", false, "Skip hop 5"),       // Boolean flag to skip hop 5
	}

	// Parse the flags to read values from the command-line arguments
	flag.Parse()

	// Ensure there is exactly one positional argument (the target address)
	if len(flag.Args()) != 1 {
		// If not, display the usage and exit with an error code
		flag.Usage()
		os.Exit(1)
	}

	// The first positional argument is the target address for traceroute
	address := flag.Args()[0]

	// Print the starting message for the traceroute
	fmt.Printf("Traceroute to %s:\n", address)

	// Call the Traceroute function from the traceroute package
	// Pass the target address and the configuration values
	err := traceroute.Traceroute(address, config)

	// If there is an error during the traceroute, print the error message and exit
	if err != nil {
		fmt.Printf("Traceroute to %s failed: %v\n", address, err)
		os.Exit(1)
	}
}
