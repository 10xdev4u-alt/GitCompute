# 🏁 Getting Started with GitCompute

Welcome to the kitchen! 🥘 This guide will take you from "zero" to "running code on GitHub's cloud" in 5 minutes.

## Prerequisites

Before we start, ensure you have:
1.  **Git** installed.
2.  **Go** (1.21+) installed (if building from source).
3.  A **GitHub Account**.

## Step 1: Get the Tool

If you haven't installed it yet:

```bash
go install github.com/10xdev4u-alt/cicd-as-a-service@latest
```

Verify it works:
```bash
git-compute --help
```

## Step 2: Prepare Your Repo

GitCompute needs a "base of operations". 

1.  Create a new GitHub repository (or use an existing one). Let's call it `my-compute-node`.
2.  Clone it locally:
    ```bash
    git clone https://github.com/YOUR_USER/my-compute-node.git
    cd my-compute-node
    ```
3.  **Crucial Step**: You need to add the Worker Workflow.
    Create `.github/workflows/worker.yml`:
    
    ```yaml
    name: GitCompute Worker
    on:
      workflow_dispatch:
        inputs:
          command:
            description: 'Command'
            required: true
          os:
            description: 'OS'
            required: true
            default: 'ubuntu-latest'
    jobs:
      compute:
        runs-on: ${{ github.event.inputs.os }}
        steps:
          - uses: actions/checkout@v3
          - run: ${{ github.event.inputs.command }} > output.log 2>&1
          - uses: actions/upload-artifact@v4
            if: always()
            with:
              name: execution-result
              path: output.log
    ```
4.  Push this to GitHub:
    ```bash
    git add .
    git commit -m "setup: add worker"
    git push origin main
    ```

## Step 3: Connect the CLI

Run the interactive setup:
```bash
git-compute init
```

- **Token**: [Generate a Classic Token here](https://github.com/settings/tokens). Select the `repo` scope.
- **Owner**: Your GitHub username.
- **Repo**: `my-compute-node` (the one you just made).

## Step 4: First Run

Let's verify everything is wired up.

```bash
git-compute run --cmd "echo Hello World"
```

**What happens?**
1.  The CLI sends a signal.
2.  You'll see "Waiting for run to start...".
3.  Then "Status: queued" -> "in_progress".
4.  Finally "Run completed!" and it prints "Hello World".

## Next Steps
Check out [Cross Compilation Examples](../examples/CROSS_COMPILATION.md) to see real power!
