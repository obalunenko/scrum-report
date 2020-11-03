package main

import "fmt"

const unset = "unset"

var ( // build info
	version     = unset
	date        = unset
	commit      = unset
	committedBy = unset
	builtBy     = unset
	goVersion   = unset
)

func printVersion() {
	fmt.Printf("version info: %s \n", version)
	fmt.Printf("Build date: %s \n", date)
	fmt.Printf("commit: %s \n", commit)
	fmt.Printf("committed by: %s \n", committedBy)
	fmt.Printf("built by: %s \n", builtBy)
	fmt.Printf("go verion: %s \n", goVersion)
	fmt.Println()
}
