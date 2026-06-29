# Pokedex CLI

A command-line Pokedex built in Go. Interact with the world of Pokemon directly from your terminal.

## Features

- Interactive REPL with a `Pokedex >` prompt
- Case-insensitive command parsing

## Getting Started

### Prerequisites

- Go 1.21+

### Install & Run

```bash
git clone https://github.com/jnguye580/pokedexcli.git
cd pokedexcli
go run .
```

### Usage

```bash
./pokedex
```

### Available Commands

| Command   | Description                    |
|-----------|--------------------------------|
| `help`    | Displays a help message        |
| `exit`    | Exit the Pokedex               |
| `explore` | Explore a location area        |
| `catch`   | Attempt to catch a Pokemon     |
| `inspect` | Inspect a caught Pokemon       |
| `pokedex` | List all caught Pokemon        |

## Requirements

- Go 1.21+

## Development

Run tests:

```bash
go test ./...
```

## Built With

- [Go](https://golang.org/)
- [PokeAPI](https://pokeapi.co/)


