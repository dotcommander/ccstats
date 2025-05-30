# ccstats

A command-line tool to analyze [Claude Code](https://claude.ai/code) usage statistics from local JSONL files.

## Overview

`ccstats` parses Claude Code API usage data from `~/.claude/projects/**/*.jsonl` files and provides detailed usage reports including:

- Token usage (input, output, cache)
- Cost analysis in USD
- Daily and session-based aggregation
- Date range filtering
- JSON output for programmatic access

## Installation

### Using Go

```bash
go install github.com/dotcommander/ccstats@latest
```

### From Source

```bash
git clone https://github.com/dotcommander/ccstats.git
cd ccstats
go build -o ccstats
```

## Usage

### Daily Usage Report

Show aggregated usage statistics by day:

```bash
ccstats daily
```

Example output:
```
Date        Input  Output  Cache Create  Cache Read  Total Tokens  Cost (USD)
--------------------------------------------------------------------------------
2025-05-20  7273   817666  7073572       267845636   275744147     $119.1664
2025-05-21  8278   520463  4147629       101390752   106067122     $53.5373
2025-05-22  17127  1208473 5556737       225422356   232204693     $216.5480
--------------------------------------------------------------------------------
TOTAL       32678  2546602 16777938      594658744   614023962     $389.2517
```

### Session Usage Report

Show usage grouped by project and session:

```bash
ccstats session
```

Example output:
```
Project     Session              Input  Output  Cache Create  Cache Read  Total Tokens  Cost (USD)  Last Activity
------------------------------------------------------------------------------------------------------------------------
my-project  session-123          6000   8000    550           425         14975         $0.4140     2025-01-30 11:20
api-server  feature-branch       12500  45000   1200          8900        67600         $1.2350     2025-01-30 15:45
------------------------------------------------------------------------------------------------------------------------
TOTAL                            18500  53000   1750          9325        82575         $1.6490
```

### Options

- `--path <dir>` - Custom path to Claude projects directory (default: `~/.claude/projects`)
- `--json` - Output in JSON format instead of table
- `--since YYYYMMDD` - Filter records since this date
- `--until YYYYMMDD` - Filter records until this date

### Examples

```bash
# Show daily usage for the current month
ccstats daily --since 20250501

# Get session report in JSON format
ccstats session --json

# Analyze specific date range
ccstats daily --since 20250101 --until 20250131

# Use custom data path
ccstats session --path /custom/path/to/.claude/projects

# Redirect errors to see clean output
ccstats daily 2>/dev/null
```

## Data Source

The tool reads JSONL files from Claude Code projects with the following structure:
```
~/.claude/projects/{project}/{session}/*.jsonl
```

Each JSONL file contains entries with:
- Timestamp
- Token usage (input, output, cache creation, cache read)
- Cost in USD
- Message type and metadata

## Output Formats

### Table Format (Default)

Human-readable tables with:
- Aligned columns
- Summary totals
- Cost in USD with 4 decimal precision
- Local timezone for timestamps

### JSON Format

Structured JSON output for integration with other tools:

```json
[
  {
    "Date": "2025-05-20",
    "InputTokens": 7273,
    "OutputTokens": 817666,
    "CacheCreateTokens": 7073572,
    "CacheReadTokens": 267845636,
    "TotalTokens": 275744147,
    "TotalCost": 119.1664
  }
]
```

## Cost Calculation

The tool uses the pre-calculated costs from Claude's API (`costUSD` field in JSONL), which already accounts for:
- Different model pricing (Opus 4, Sonnet 3.5, Haiku, etc.)
- Cache read discount (10% of normal input token price)
- Cache creation costs (25% of normal input token price)
- Standard input/output token rates per model

This ensures accurate cost reporting regardless of which Claude model was used.

## Features

- **Fast Performance**: Streams files line-by-line to minimize memory usage
- **Large File Support**: Handles JSONL lines up to 10MB
- **Error Resilience**: Continues processing when encountering malformed data
- **Cross-Platform**: Works on macOS, Linux, and Windows
- **Zero Dependencies**: Only requires Go standard library and Cobra for CLI

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

Created for the Claude Code community to better understand their API usage and costs.