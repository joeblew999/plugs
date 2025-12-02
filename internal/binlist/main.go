// Package binlist is a dev tool that prints plugin names from plugins.json.
// Used by Taskfile to derive the BINARIES variable.
// This is NOT a released binary - it lives in internal/, not cmd/.
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joeblew999/plugs/internal/registry"
)

func main() {
	reg, err := registry.LoadFromRoot()
	if err != nil {
		fmt.Fprintf(os.Stderr, "binlist: %v\n", err)
		os.Exit(1)
	}

	names := reg.Binaries()
	fmt.Println(strings.Join(names, " "))
}
