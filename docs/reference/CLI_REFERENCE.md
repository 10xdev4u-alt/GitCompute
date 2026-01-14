# 📖 CLI Reference Manual

This document details every flag and command available in `git-compute`.

## Global Flags
These apply to all commands.

- `--config string`: Path to the config file. (Default: `$HOME/.git-compute.yaml`)

## Commands

### `init`
Initializes the configuration.

**Usage:**
```bash
git-compute init
```
**Interactive Prompts:**
- **GitHub Token**: A Classic PAT with `repo` scope.
- **Owner**: Repository owner.
- **Repo**: Repository name.

---

### `run`
Dispatches a job to the remote runner.

**Usage:**
```bash
git-compute run [flags]
```

**Flags:**

| Flag | Shorthand | Type | Default | Description |
|------|-----------|------|---------|-------------|
| `--cmd` | `-c` | `string` | `"echo Hello World"` | The actual shell command to run on the VM. |
| `--os` | `-o` | `string` | `"ubuntu-latest"` | The target OS. Options: `ubuntu-latest`, `windows-latest`, `macos-latest`. |
| `--watch` | `-w` | `bool` | `true` | If true, blocks and polls for status. If false, triggers and exits. |

**Examples:**

*Run silently (fire and forget):*
```bash
git-compute run --cmd "long_script.sh" --watch=false
```

*Run on Windows:*
```bash
git-compute run --cmd "dir C:\\" --os windows-latest
```

