# ⚙️ Internal Protocol

This document describes how the CLI talks to the Worker. Useful for contributors.

## 1. The Trigger (`workflow_dispatch`)

The CLI makes a POST request to:
`https://api.github.com/repos/{owner}/{repo}/actions/workflows/{id}/dispatches`

**Payload:**
```json
{
  "ref": "main",
  "inputs": {
    "command": "echo Hello",
    "os": "ubuntu-latest"
  }
}
```

## 2. The Execution (Worker)

The Worker is a standard GitHub Action.
1. It reads `github.event.inputs.command`.
2. It executes it via `bash` (or `powershell` on Windows).
3. It redirects stdout/stderr to `output.log`.
4. It uses `actions/upload-artifact` to push `output.log`.

## 3. The Retrieval (Polling)

1. CLI lists runs for the workflow.
2. Finds the most recent run (created < 1 min ago).
3. Polls `GET /repos/.../runs/{run_id}` until `status == completed`.
4. Lists artifacts for the run.
5. Downloads the artifact blob (zip).
6. Unzips locally.
