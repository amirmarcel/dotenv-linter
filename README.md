# dotenv-linter
![CI](https://github.com/amirmarcel/dotenv-linter/actions/workflows/ci.yml/badge.svg)
A CLI tool that diffs `.env` files against `.env.example` to catch missing or extra keys — with CI-friendly exit codes.

## The Problem

When a new key is added to `.env.example` but a developer forgets to add it to their local `.env`, the app breaks at runtime — not at boot. This tool catches that drift before it becomes a bug.

## Usage

```bash
dotenv-linter --env .env --example .env.example
```

## Output

```
x Missing in .env:         SECRET_KEY
x Missing in .env.example: DEBUG_MODE
```

Exits `0` if all keys are in sync. Exits `1` if drift is detected — making it suitable for pre-commit hooks and CI pipelines.

## Installation

```bash
go install github.com/amirmarcel/dotenv-linter@latest
```

## Flags

| Flag        | Default          | Description              |
|-------------|------------------|--------------------------|
| `--env`     | `.env`           | Path to your local .env  |
| `--example` | `.env.example`   | Path to your template    |

## Built With

- Go standard library only — no dependencies