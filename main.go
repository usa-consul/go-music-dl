// go-music-dl - A music downloader written in Go
// Fork of guohuiyuan/go-music-dl
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/misitebao/go-music-dl/cmd"
)

var (
	// Version information injected at build time via ldflags
	Version   = "dev"
	Commit    = "none"
	BuildDate = "unknown"
)

func main() {
	// Define top-level flags
	versionFlag := flag.Bool("version", false, "Print version information and exit")
	vFlag := flag.Bool("v", false, "Print version information and exit (shorthand)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "go-music-dl - A music downloader\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n  go-music-dl [options] <command> [arguments]\n\n")
		fmt.Fprintf(os.Stderr, "Commands:\n")
		fmt.Fprintf(os.Stderr, "  download\tDownload music from supported platforms\n")
		fmt.Fprintf(os.Stderr, "  search\t\tSearch for music across platforms\n")
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *versionFlag || *vFlag {
		printVersion()
		os.Exit(0)
	}

	// Execute the root command
	if err := cmd.Execute(Version, Commit, BuildDate); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// printVersion outputs build version information to stdout
func printVersion() {
	fmt.Printf("go-music-dl (personal fork)\n")
	fmt.Printf("  Version:    %s\n", Version)
	fmt.Printf("  Commit:     %s\n", Commit)
	fmt.Printf("  Build Date: %s\n", BuildDate)
	fmt.Printf("  Upstream:   https://github.com/guohuiyuan/go-music-dl\n")
}
