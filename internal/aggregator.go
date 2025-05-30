package internal

import (
	"sort"
	"time"
)

// AggregateDailyUsage groups usage records by day
func AggregateDailyUsage(records []UsageRecord) []DailySummary {
	dailyMap := make(map[string]*DailySummary)
	
	for _, record := range records {
		// Use local date for grouping
		dateStr := record.Timestamp.Local().Format("2006-01-02")
		
		if summary, exists := dailyMap[dateStr]; exists {
			summary.InputTokens += record.InputTokens
			summary.OutputTokens += record.OutputTokens
			summary.CacheCreateTokens += record.CacheCreateTokens
			summary.CacheReadTokens += record.CacheReadTokens
			summary.TotalCost += record.CostUSD
		} else {
			dailyMap[dateStr] = &DailySummary{
				Date:              dateStr,
				InputTokens:       record.InputTokens,
				OutputTokens:      record.OutputTokens,
				CacheCreateTokens: record.CacheCreateTokens,
				CacheReadTokens:   record.CacheReadTokens,
				TotalCost:         record.CostUSD,
			}
		}
	}
	
	// Convert map to slice and calculate total tokens
	var summaries []DailySummary
	for _, summary := range dailyMap {
		summary.TotalTokens = summary.InputTokens + summary.OutputTokens + 
			summary.CacheCreateTokens + summary.CacheReadTokens
		summaries = append(summaries, *summary)
	}
	
	// Sort by date
	sort.Slice(summaries, func(i, j int) bool {
		return summaries[i].Date < summaries[j].Date
	})
	
	return summaries
}

// AggregateSessionUsage groups usage records by project/session
func AggregateSessionUsage(records []UsageRecord) []SessionSummary {
	sessionMap := make(map[string]*SessionSummary)
	
	for _, record := range records {
		key := record.Project + "/" + record.Session
		
		if summary, exists := sessionMap[key]; exists {
			summary.InputTokens += record.InputTokens
			summary.OutputTokens += record.OutputTokens
			summary.CacheCreateTokens += record.CacheCreateTokens
			summary.CacheReadTokens += record.CacheReadTokens
			summary.TotalCost += record.CostUSD
			
			// Update last activity if this record is more recent
			if record.Timestamp.After(summary.LastActivity) {
				summary.LastActivity = record.Timestamp
			}
		} else {
			sessionMap[key] = &SessionSummary{
				Project:           record.Project,
				Session:           record.Session,
				InputTokens:       record.InputTokens,
				OutputTokens:      record.OutputTokens,
				CacheCreateTokens: record.CacheCreateTokens,
				CacheReadTokens:   record.CacheReadTokens,
				TotalCost:         record.CostUSD,
				LastActivity:      record.Timestamp,
			}
		}
	}
	
	// Convert map to slice and calculate total tokens
	var summaries []SessionSummary
	for _, summary := range sessionMap {
		summary.TotalTokens = summary.InputTokens + summary.OutputTokens + 
			summary.CacheCreateTokens + summary.CacheReadTokens
		summaries = append(summaries, *summary)
	}
	
	// Sort by last activity (most recent first)
	sort.Slice(summaries, func(i, j int) bool {
		return summaries[i].LastActivity.After(summaries[j].LastActivity)
	})
	
	return summaries
}

// FilterRecordsByDate filters records based on since/until dates
func FilterRecordsByDate(records []UsageRecord, since, until time.Time) []UsageRecord {
	if since.IsZero() && until.IsZero() {
		return records
	}
	
	var filtered []UsageRecord
	for _, record := range records {
		if !since.IsZero() && record.Timestamp.Before(since) {
			continue
		}
		if !until.IsZero() && record.Timestamp.After(until) {
			continue
		}
		filtered = append(filtered, record)
	}
	
	return filtered
}