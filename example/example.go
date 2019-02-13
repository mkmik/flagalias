package main

import (
	"flag"
	"fmt"

	"github.com/mkmik/flagalias"
)

var (
	dryRun = flag.Bool("dry-run", false, "dry run")
)

func init() {
	flagalias.Alias("N", "dry-run")
	flagalias.Deprecated("dryrun", "dry-run")
}

func main() {
	flag.Parse()

	fmt.Printf("dry run: %v\n", *dryRun)
}
