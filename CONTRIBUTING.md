# 🤝 Contributing to GitCompute

First off, thank you for considering contributing to GitCompute (ActionGrid)! We are building the "Poor Man's Supercomputer" and every bit of help counts.

## 🛠️ Development Setup

1. **Fork & Clone**
   ```bash
   git clone https://github.com/YOUR_USERNAME/cicd-as-a-service.git
   cd cicd-as-a-service
   ```

2. **Install Dependencies**
   ```bash
   go mod download
   ```

3. **Run Tests**
   ```bash
   go test ./...
   ```

## 📐 Coding Standards

- **Language**: Go (Golang) 1.21+
- **Style**: Standard `gofmt`. Please run `go fmt ./...` before committing.
- **Linting**: We recommend `golangci-lint`.
- **Architecture**: Follow the `cmd/` (CLI) and `pkg/` (Library) separation.

## 📝 Commit Messages

We strictly follow [Conventional Commits](https://www.conventionalcommits.org/).

**Format**: `<type>(<scope>): <subject>`

**Types**:
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation only changes
- `style`: Changes that do not affect the meaning of the code (white-space, formatting, etc)
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `perf`: A code change that improves performance
- `test`: Adding missing tests or correcting existing tests
- `chore`: Changes to the build process or auxiliary tools

**Example**:
```text
feat(runner): implement artifact downloading logic
fix(cli): correct flag parsing for --os
docs: update architecture diagrams
```

## 🚀 Pull Request Process

1. Create a new branch: `git checkout -b feature/amazing-feature`.
2. Commit your changes (remember the commit message standards!).
3. Push to the branch: `git push origin feature/amazing-feature`.
4. Open a Pull Request.
5. Wait for the review from **PrinceTheProgrammer** or the core team.

## 🐛 Reporting Bugs

Please open an issue on GitHub with:
1. The command you ran.
2. The expected output.
3. The actual output (and `output.log` if available).
