package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	
	"github.com/spf13/cobra"
)

var (
	dataPath string
	jsonOutput bool
	since string
	until string
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "ccstats",
	Short: "Analyze Claude Code usage from local JSONL files",
	Long: `ccstats is a CLI tool to analyze Claude Code usage from local JSONL files.
	
It parses and aggregates usage data from ~/.claude/projects/**/*.jsonl files,
providing daily and session-based usage/cost reports with table or JSON output.`,
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Get default data path
	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	}
	defaultPath := filepath.Join(home, ".claude", "projects")
	
	// Global flags
	rootCmd.PersistentFlags().StringVar(&dataPath, "path", defaultPath, "Path to Claude projects directory")
	rootCmd.PersistentFlags().BoolVar(&jsonOutput, "json", false, "Output in JSON format")
	rootCmd.PersistentFlags().StringVar(&since, "since", "", "Filter records since this date (YYYYMMDD)")
	rootCmd.PersistentFlags().StringVar(&until, "until", "", "Filter records until this date (YYYYMMDD)")
}