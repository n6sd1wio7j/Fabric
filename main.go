package main

import (
	"os"

	"github.com/danielmiessler/fabric/cli"
	"github.com/danielmiessler/fabric/core"
)

// main is the entry point for the Fabric CLI application.
// Fabric is an open-source AI augmentation framework that helps humans
// apply AI to everyday tasks through a modular pattern-based approach.
func main() {
	// Initialize the Fabric registry which manages all available patterns,
	// models, and integrations.
	registry, err := core.NewRegistry()
	if err != nil {
		println("Error initializing Fabric registry:", err.Error())
		os.Exit(1)
	}

	// Run the CLI with the initialized registry.
	// The CLI handles argument parsing, pattern selection, and output.
	if err := cli.Run(registry); err != nil {
		println("Error:", err.Error())
		os.Exit(1)
	}
}
