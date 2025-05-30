package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// OutputDailyTable prints daily summaries in a table format
func OutputDailyTable(summaries []DailySummary) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	
	// Header
	fmt.Fprintln(w, "Date\tInput\tOutput\tCache Create\tCache Read\tTotal Tokens\tCost (USD)")
	fmt.Fprintln(w, strings.Repeat("-", 80))
	
	var totalInput, totalOutput, totalCacheCreate, totalCacheRead, totalTokens int
	var totalCost float64
	
	for _, summary := range summaries {
		fmt.Fprintf(w, "%s\t%d\t%d\t%d\t%d\t%d\t$%.4f\n",
			summary.Date,
			summary.InputTokens,
			summary.OutputTokens,
			summary.CacheCreateTokens,
			summary.CacheReadTokens,
			summary.TotalTokens,
			summary.TotalCost,
		)
		
		totalInput += summary.InputTokens
		totalOutput += summary.OutputTokens
		totalCacheCreate += summary.CacheCreateTokens
		totalCacheRead += summary.CacheReadTokens
		totalTokens += summary.TotalTokens
		totalCost += summary.TotalCost
	}
	
	// Total row
	fmt.Fprintln(w, strings.Repeat("-", 80))
	fmt.Fprintf(w, "TOTAL\t%d\t%d\t%d\t%d\t%d\t$%.4f\n",
		totalInput, totalOutput, totalCacheCreate, totalCacheRead, totalTokens, totalCost)
	
	w.Flush()
}

// OutputSessionTable prints session summaries in a table format
func OutputSessionTable(summaries []SessionSummary) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	
	// Header
	fmt.Fprintln(w, "Project\tSession\tInput\tOutput\tCache Create\tCache Read\tTotal Tokens\tCost (USD)\tLast Activity")
	fmt.Fprintln(w, strings.Repeat("-", 120))
	
	var totalInput, totalOutput, totalCacheCreate, totalCacheRead, totalTokens int
	var totalCost float64
	
	for _, summary := range summaries {
		fmt.Fprintf(w, "%s\t%s\t%d\t%d\t%d\t%d\t%d\t$%.4f\t%s\n",
			summary.Project,
			truncateString(summary.Session, 20),
			summary.InputTokens,
			summary.OutputTokens,
			summary.CacheCreateTokens,
			summary.CacheReadTokens,
			summary.TotalTokens,
			summary.TotalCost,
			summary.LastActivity.Local().Format("2006-01-02 15:04"),
		)
		
		totalInput += summary.InputTokens
		totalOutput += summary.OutputTokens
		totalCacheCreate += summary.CacheCreateTokens
		totalCacheRead += summary.CacheReadTokens
		totalTokens += summary.TotalTokens
		totalCost += summary.TotalCost
	}
	
	// Total row
	fmt.Fprintln(w, strings.Repeat("-", 120))
	fmt.Fprintf(w, "TOTAL\t\t%d\t%d\t%d\t%d\t%d\t$%.4f\t\n",
		totalInput, totalOutput, totalCacheCreate, totalCacheRead, totalTokens, totalCost)
	
	w.Flush()
}

// OutputDailyJSON prints daily summaries as JSON
func OutputDailyJSON(summaries []DailySummary) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(summaries)
}

// OutputSessionJSON prints session summaries as JSON
func OutputSessionJSON(summaries []SessionSummary) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(summaries)
}

// truncateString truncates a string to a maximum length
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
}