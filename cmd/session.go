package cmd

import (
	"fmt"
	"os"
	
	"github.com/dotcommander/ccstats/internal"
	"github.com/spf13/cobra"
)

// sessionCmd represents the session command
var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Aggregate and display usage/cost by session",
	Long:  `Show session-based usage statistics including token counts and costs, grouped by project and session.`,
	RunE:  runSession,
}

func init() {
	rootCmd.AddCommand(sessionCmd)
}

func runSession(cmd *cobra.Command, args []string) error {
	// Parse date filters
	sinceTime, err := parseDate(since)
	if err != nil {
		return fmt.Errorf("invalid since date: %w", err)
	}
	
	untilTime, err := parseDate(until)
	if err != nil {
		return fmt.Errorf("invalid until date: %w", err)
	}
	
	// Load records
	records, err := internal.LoadUsageRecords(dataPath, sinceTime, untilTime)
	if err != nil {
		return fmt.Errorf("failed to load usage records: %w", err)
	}
	
	if len(records) == 0 {
		fmt.Fprintln(os.Stderr, "No usage records found")
		return nil
	}
	
	// Aggregate by session
	summaries := internal.AggregateSessionUsage(records)
	
	// Output results
	if jsonOutput {
		return internal.OutputSessionJSON(summaries)
	}
	
	internal.OutputSessionTable(summaries)
	return nil
}