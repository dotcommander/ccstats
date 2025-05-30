package internal

import "time"

// UsageRecord represents a single usage entry from JSONL files
type UsageRecord struct {
	Timestamp         time.Time
	Project           string
	Session           string
	InputTokens       int
	OutputTokens      int
	CacheCreateTokens int
	CacheReadTokens   int
	CostUSD           float64
}

// DailySummary aggregates usage data by day
type DailySummary struct {
	Date              string
	InputTokens       int
	OutputTokens      int
	CacheCreateTokens int
	CacheReadTokens   int
	TotalTokens       int
	TotalCost         float64
}

// SessionSummary aggregates usage data by project/session
type SessionSummary struct {
	Project           string
	Session           string
	InputTokens       int
	OutputTokens      int
	CacheCreateTokens int
	CacheReadTokens   int
	TotalTokens       int
	TotalCost         float64
	LastActivity      time.Time
}