package cmd

import (
	"fmt"
	"os"
	"time"
	
	"github.com/dotcommander/ccstats/internal"
	"github.com/spf13/cobra"
)

// dailyCmd represents the daily command
var dailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "Aggregate and display usage/cost by day",
	Long:  `Show daily usage statistics including token counts and costs.`,
	RunE:  runDaily,
}

func init() {
	rootCmd.AddCommand(dailyCmd)
}

func runDaily(cmd *cobra.Command, args []string) error {
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
	
	// Aggregate by day
	summaries := internal.AggregateDailyUsage(records)
	
	// Output results
	if jsonOutput {
		return internal.OutputDailyJSON(summaries)
	}
	
	internal.OutputDailyTable(summaries)
	return nil
}

// parseDate parses a date string in YYYYMMDD format
func parseDate(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, nil
	}
	
	return time.Parse("20060102", dateStr)
}