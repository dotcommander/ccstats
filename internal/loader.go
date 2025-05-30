package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// JSONLEntry represents the structure of each JSONL line
type JSONLEntry struct {
	Type      string    `json:"type"`
	Timestamp string    `json:"timestamp"`
	CostUSD   float64   `json:"costUSD"`
	Message   *Message  `json:"message"`
}

// Message contains the usage information
type Message struct {
	Usage *Usage `json:"usage"`
}

// Usage represents token usage information
type Usage struct {
	InputTokens         int `json:"input_tokens"`
	OutputTokens        int `json:"output_tokens"`
	CacheCreationTokens int `json:"cache_creation_input_tokens"`
	CacheReadTokens     int `json:"cache_read_input_tokens"`
}

// LoadUsageRecords loads all usage records from JSONL files in the specified directory
func LoadUsageRecords(dataPath string, since, until time.Time) ([]UsageRecord, error) {
	var records []UsageRecord
	
	err := filepath.Walk(dataPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if !strings.HasSuffix(path, ".jsonl") {
			return nil
		}
		
		fileRecords, err := loadJSONLFile(path, since, until)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: error loading %s: %v\n", path, err)
			return nil // Skip file but continue
		}
		
		records = append(records, fileRecords...)
		return nil
	})
	
	return records, err
}

// loadJSONLFile loads records from a single JSONL file
func loadJSONLFile(path string, since, until time.Time) ([]UsageRecord, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	// Extract project and session from path
	// Expected format: ~/.claude/projects/{project}/{session}/api_requests.jsonl
	parts := strings.Split(filepath.ToSlash(path), "/")
	var project, session string
	
	for i, part := range parts {
		if part == "projects" && i+2 < len(parts) {
			project = parts[i+1]
			session = parts[i+2]
			break
		}
	}
	
	var records []UsageRecord
	scanner := bufio.NewScanner(file)
	// Increase buffer size to handle large JSONL lines (up to 10MB)
	const maxScanTokenSize = 10 * 1024 * 1024
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, maxScanTokenSize)
	lineNum := 0
	
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		
		var entry JSONLEntry
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: skipping malformed JSON at %s:%d: %v\n", path, lineNum, err)
			continue
		}
		
		// Skip entries that aren't assistant messages with usage data
		if entry.Type != "assistant" || entry.Message == nil || entry.Message.Usage == nil {
			continue
		}
		
		// Skip if timestamp is empty
		if entry.Timestamp == "" {
			continue
		}
		
		// Parse timestamp
		timestamp, err := time.Parse(time.RFC3339, entry.Timestamp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: invalid timestamp at %s:%d: %v\n", path, lineNum, err)
			continue
		}
		
		// Apply date filters
		if !since.IsZero() && timestamp.Before(since) {
			continue
		}
		if !until.IsZero() && timestamp.After(until) {
			continue
		}
		
		record := UsageRecord{
			Timestamp:         timestamp,
			Project:           project,
			Session:           session,
			InputTokens:       entry.Message.Usage.InputTokens,
			OutputTokens:      entry.Message.Usage.OutputTokens,
			CacheCreateTokens: entry.Message.Usage.CacheCreationTokens,
			CacheReadTokens:   entry.Message.Usage.CacheReadTokens,
			CostUSD:           entry.CostUSD,
		}
		
		records = append(records, record)
	}
	
	if err := scanner.Err(); err != nil && err != io.EOF {
		return records, fmt.Errorf("error reading file: %w", err)
	}
	
	return records, nil
}