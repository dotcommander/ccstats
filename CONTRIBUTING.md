# Contributing to ccstats

Thank you for your interest in contributing to ccstats! We welcome contributions from the community.

## How to Contribute

1. **Fork the repository** and create your branch from `main`.
2. **Make your changes** and ensure the code follows Go conventions.
3. **Test your changes** thoroughly.
4. **Submit a pull request** with a clear description of your changes.

## Development Setup

```bash
# Clone your fork
git clone https://github.com/yourusername/ccstats.git
cd ccstats

# Install dependencies
go mod download

# Build the project
go build

# Run tests
go test ./...
```

## Code Style

- Follow standard Go formatting (use `gofmt`)
- Write clear, self-documenting code
- Add comments for complex logic
- Keep functions focused and small

## Testing

- Add tests for new functionality
- Ensure all tests pass before submitting PR
- Test with real Claude Code JSONL files if possible

## Reporting Issues

- Use GitHub Issues to report bugs
- Include steps to reproduce the issue
- Provide example JSONL data if relevant (sanitized)

## Feature Requests

- Open an issue to discuss new features
- Explain the use case and benefits
- Consider implementing it yourself!

Thank you for contributing!