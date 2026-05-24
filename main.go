package main

import (
	"fmt"
	"os"

	"github.com/danielmiessler/fabric/cli"
	"github.com/danielmiessler/fabric/core"
)

// main is the entry point for the Fabric CLI application.
// Fabric is an open-source AI augmentation framework that helps humans
// apply AI to everyday tasks through a modular pattern-based approach.
//
// Personal fork notes:
//   - Using this for local experiments with custom patterns
//   - See ./patterns/local/ for my personal pattern additions
//   - Fork maintained at: github.com/myusername/fabric
//   - Upstream: github.com/danielmiessler/fabric
//
// TODO: explore adding a --dry-run flag to preview pattern substitutions
//       without sending requests to the model
// TODO: look into adding shell completion support (bash/zsh)
func main() {
	// Initialize the Fabric registry which manages all available patterns,
	// models, and integrations.
	registry, err := core.NewRegistry()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing Fabric registry: %v\n", err)
		os.Exit(1)
	}

	// Run the CLI with the initialized registry.
	// The CLI handles argument parsing, pattern selection, and output.
	if err := cli.Run(registry); err != nil {
		// Print a more descriptive prefix to make errors easier to spot in logs.
		// Exit code 2 used here to distinguish CLI/usage errors from init errors (1).
		fmt.Fprintf(os.Stderr, "[fabric] Error: %v\n", err)
		os.Exit(2)
	}
}
