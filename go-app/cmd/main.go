package main

import (
	"fmt"
	"os"
	"runtime"
)

type sysInfo struct {
	Hostname string
	OS       string
	Arch     string
	GoVer    string
}

func gatherInfo() sysInfo {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	return sysInfo{
		Hostname: hostname,
		OS:       runtime.GOOS,
		Arch:     runtime.GOARCH,
		GoVer:    runtime.Version(),
	}
}

func main() {
	info := gatherInfo()
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║       Hello from Go + Docker!    ║")
	fmt.Println("╚══════════════════════════════════╝")
	fmt.Printf("  Hostname : %s\n", info.Hostname)
	fmt.Printf("  OS       : %s\n", info.OS)
	fmt.Printf("  Arch     : %s\n", info.Arch)
	fmt.Printf("  Go ver   : %s\n", info.GoVer)
}
