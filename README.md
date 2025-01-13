I know ... I am awesome.

## Next steps

- [ ] Add tool support and verify if it works
- [ ] Add more tool templates
- [ ] Add support for init

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
./restack init <project_name>
./restack workflow <name>
./restack function <name>
```

## Example

```bash
./restack workflow query_companies
./restack function query_companies
```
