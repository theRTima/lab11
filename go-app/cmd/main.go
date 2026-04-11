package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║       Hello from Go + Docker!    ║")
	fmt.Println("╚══════════════════════════════════╝")
	fmt.Printf("  Hostname : %s\n", hostname)
	fmt.Printf("  OS       : %s\n", runtime.GOOS)
	fmt.Printf("  Arch     : %s\n", runtime.GOARCH)
	fmt.Printf("  Go ver   : %s\n", runtime.Version())
}
