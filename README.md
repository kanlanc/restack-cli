I know ... I am awesome.

## From source

Download the source code and build the CLI.

```bash
go build -o restack cmd/cli/main.go
```

## Installation

```bash
# go install github.com/kanlanc/restack-cli/cmd/cli
```

## Usage

```bash
./restack workflow <name>
./restack function <name>
```

## Example

```bash
./restack workflow query_companies
./restack function query_companies
```
