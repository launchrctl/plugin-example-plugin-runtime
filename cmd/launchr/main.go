// Package executes Launchr application.
package main

import (
	"github.com/launchrctl/launchr"

	_ "github.com/launchrctl/plugin-example-plugin-runtime"
)

func main() {
	launchr.RunAndExit()
}
