# 📖 GitCompute Usage Guide

## ⚙️ Setup

### 1. Installation
Build from source:
```bash
git clone https://github.com/10xdev4u-alt/cicd-as-a-service.git
cd cicd-as-a-service
go install .
```

### 2. Configuration
Run the interactive initializer:
```bash
git-compute init
```
You will be asked for:
- **GitHub Token**: Generate one [here](https://github.com/settings/tokens) with `repo` scope.
- **Owner**: Your GitHub username or organization (e.g., `10xdev4u-alt`).
- **Repo**: The repository where `worker.yml` is installed.

### 3. Deploy Worker
Ensure the workflow is present in your target repository:
```bash
# Push the workflow to GitHub
git add .github/workflows/worker.yml
git commit -m "feat: add git-compute worker"
git push origin main
```

---

## 💻 Commands

### `run` - Execute Remote Command
Triggers a job on a GitHub Runner.

#### Syntax
```bash
git-compute run [flags]
```

#### Flags
| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--cmd` | `-c` | The shell command to execute. | `echo Hello World` |
| `--os` | `-o` | OS Image (`ubuntu-latest`, `windows-latest`, `macos-latest`). | `ubuntu-latest` |
| `--watch` | `-w` | Watch status and download logs automatically. | `true` |

#### Examples

**1. Simple Hello World (Linux)**
```bash
git-compute run --cmd "echo Hello from GitCompute"
```

**2. System Info (Windows)**
```bash
git-compute run --cmd "systeminfo" --os windows-latest
```

**3. Compile Go Binary (Cross-Compilation)**
```bash
git-compute run --cmd "go build -o myapp.exe main.go" --os windows-latest
```

**4. Run Python Script**
```bash
git-compute run --cmd "python3 -c 'print(sum(range(100)))'"
```

---

## 📂 Artifacts
When a job completes successfully:
1. The CLI downloads a zip file (e.g., `result-12345.zip`).
2. It automatically extracts it to your current directory.
3. The standard output/error is saved in `output.log`.
