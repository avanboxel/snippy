# Snippy

A simple offline snippet manager for code snippets.

## Installation

```bash
go install github.com/avanboxel/snippy@latest
```

Or build from source:

```bash
git clone https://github.com/avanboxel/snippy.git
cd snippy
go build -o snippy
```

## Usage

### Add a snippet

```bash
# Add from command line
snippy add "fmt.Println(\"Hello World\")" --lang go --tags example,basic

# Add from stdin
echo "console.log('Hello World')" | snippy add --lang js --tags example

# Add from file
snippy add --lang go --tags example < app.js
```

### List snippets

```bash
# List all snippets
snippy list

# Filter by language
snippy list --lang go
snippy list -l go

# Filter by tags
snippy list --tags example
snippy list -t example

# Search by part of code
snippy list --search "Hello World"
snippy list -s "Hello World"
```

### Clean snippets

```bash
# Clean all snippets
snippy clean

# Clean snippet by id
snippy clean --id 1234
snippy clean -i 1234

# Clean all snippets by language
snippy clean --lang go
snippy clean -l go

# Clean all snippets by tags
snippy clean --tags example
snippy clean -t example

# Clean all snippets by part of code
snippy clean --search "Hello World"
snippy clean -s "Hello World"
```

## License

MIT