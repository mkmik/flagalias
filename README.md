[![GoDoc](https://godoc.org/github.com/mkmik/flagalias?status.svg)](http://godoc.org/github.com/mkmik/flagalias)

# Overview

Package flagalias allows you to define aliases for Go flags.
Deprecated aliases will emit a warning when used.

## Example

```go
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
```