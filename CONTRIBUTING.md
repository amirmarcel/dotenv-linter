# Contributing

Contributions are welcome! Here's how to get started.

## Setup

```bash
git clone https://github.com/amirmarcel/dotenv-linter.git
cd dotenv-linter
go mod tidy
```

## Running Tests

```bash
go test ./...
```

## Submitting Changes

1. Fork the repo
2. Create a branch: `git checkout -b feat/your-feature`
3. Make your changes
4. Run tests: `go test ./...`
5. Commit using conventional commits: `git commit -m "feat: your feature"`
6. Push and open a pull request

## Conventional Commits

This project uses [conventional commits](https://conventionalcommits.org):

- `feat:` - new feature
- `fix:` - bug fix
- `docs:` - documentation only
- `refactor:` - code change that isn't a fix or feature
- `ci:` - CI/CD changes
- `chore:` - maintenance`