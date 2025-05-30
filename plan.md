# ccusage (Go Implementation) — Architectural Instructions

## Overview

Rewrite the ccusage CLI tool in Go to analyze Claude Code usage from local JSONL files. The Go tool should:
- Parse and aggregate usage data from local files (~/.claude/projects/**/*.jsonl).
- Provide daily and session-based usage/cost reports.
- Support table and JSON output.
- Be easy to use, robust, and cross-platform.

---

## Main Features

- **CLI Commands:**
  - `daily`: Aggregate and display usage/cost by day.
  - `session`: Aggregate and display usage/cost by session.
  - Options for date filtering (`--since`, `--until`), custom data path, JSON output (`--json`), help/version.

- **Data Loading:**
  - Recursively scan a specified directory (default: ~/.claude/projects) for `.jsonl` files.
  - Parse each line as a JSON object, extracting timestamp, token counts, and cost.

- **Aggregation:**
  - For `daily`: Group usage by date, sum tokens and cost.
  - For `session`: Group by project/session, sum tokens and cost, track last activity date.

- **Output:**
  - Pretty tables for terminal (use a Go table library, e.g. github.com/olekukonko/tablewriter).
  - Optionally emit structured JSON for programmatic use.

- **Error Handling:**
  - Skip malformed lines with warnings.
  - Handle missing or unreadable files gracefully.

- **Testing:**
  - Unit tests for data parsing, aggregation, and command logic.

---

## Key Go Packages

- **CLI**: Use [spf13/cobra](https://github.com/spf13/cobra) for CLI commands and flags.
- **File I/O**: Standard library (`os`, `io`, `path/filepath`).
- **JSONL Parsing**: Use `bufio.Scanner` to read line-by-line and `encoding/json` for decoding.
- **Tables**: [olekukonko/tablewriter](https://github.com/olekukonko/tablewriter) for terminal tables.
- **Testing**: Standard `testing` package.

---

## Project Structure

```
ccusage-go/
  cmd/
    root.go         # CLI entry point and shared flags
    daily.go        # 'daily' command
    session.go      # 'session' command
  internal/
    loader.go       # Loads and parses data files
    aggregator.go   # Aggregates data for reports
    output.go       # Table/JSON output helpers
    types.go        # Data structures
  main.go           # Main executable
  go.mod
  go.sum
  README.md
```

---

## Data Structures

```go
type UsageRecord struct {
    Timestamp   time.Time
    Project     string
    Session     string
    InputTokens int
    OutputTokens int
    CacheCreateTokens int
    CacheReadTokens int
    CostUSD     float64
}

type DailySummary struct {
    Date         string
    InputTokens  int
    OutputTokens int
    CacheCreateTokens int
    CacheReadTokens int
    TotalTokens  int
    TotalCost    float64
}

type SessionSummary struct {
    Project     string
    Session     string
    InputTokens int
    OutputTokens int
    CacheCreateTokens int
    CacheReadTokens int
    TotalTokens int
    TotalCost   float64
    LastActivity time.Time
}
```

---

## Algorithm Sketch

1. **CLI Invoked**: Parse flags, get subcommand (`daily` or `session`).
2. **Load Data**: Recursively find all `.jsonl` files in the target directory.
3. **Parse Records**: For each file, read lines, decode JSON, build `UsageRecord` objects.
4. **Aggregate**:
    - For daily: Group records by date (local timezone).
    - For session: Group by project/session.
5. **Output**:
    - If `--json`, marshal summaries to JSON.
    - Else, pretty-print table to terminal.

---

## Key Design Considerations

- **Performance**: Stream files line-by-line to minimize memory.
- **Cross-platform**: Use `os.UserHomeDir()` and `filepath` for path handling.
- **Extensibility**: Keep loader/aggregator/output logic modular.
- **Testing**: Mock filesystem for unit tests; test all aggregation logic.
- **Error Handling**: Log and skip malformed lines, don't abort unless critical.
- **Documentation**: Document CLI usage, flags, and output formats in README.md.

---

## Example Commands

```bash
# Show daily usage report
ccusage-go daily

# Show session report (JSON output, custom data path)
ccusage-go session --json --path /custom/path/to/.claude/projects

# Filter by date
ccusage-go daily --since 20250501 --until 20250530
```

---

## Extensions (Future Work)

- Config file support for default paths/options.
- Export CSV.
- Web UI for visualization.
- Optional cloud sync.

---

## Inspiration

- Original ccusage architecture (TypeScript).
- [spf13/cobra](https://github.com/spf13/cobra) for Go CLI.
- [olekukonko/tablewriter](https://github.com/olekukonko/tablewriter) for pretty tables.

---

## References

- Original ccusage repo: https://github.com/ryoppippi/ccusage
- DuckDB usage analysis: https://note.com/milliondev/n/n1d018da2d769