# High Level Design (HLD) - GitCompute

## 1. Overview
GitCompute orchestrates distributed computing tasks using GitHub Actions as the execution layer and the GitHub API as the control plane. It decouples the compute environment from the triggering client.

## 2. System Components

### A. The Client (CLI)
- **Role**: The Controller.
- **Responsibility**: 
    - Accepts user input (commands, scripts, files).
    - Packages payload (compression/encryption).
    - Triggers the remote execution via GitHub API (`workflow_dispatch`).
    - Polls for completion.
    - Downloads and reassembles results (Artifacts).

### B. The Transport Layer (GitHub API)
- **Role**: The Message Bus.
- **Responsibility**:
    - Authenticates requests.
    - Queues jobs (Workflows).
    - Stores intermediate data (Artifacts).

### C. The Worker (GitHub Action)
- **Role**: The Compute Node.
- **Responsibility**:
    - Listens for the `workflow_dispatch` event.
    - Unpacks the payload.
    - Executes the user-defined logic (Shell, Docker, etc.).
    - Captures `stdout`/`stderr`.
    - Uploads outputs as a GitHub Artifact.

## 3. Data Flow

1. **Trigger**: User runs `git-compute run "make build"`.
2. **Dispatch**: CLI sends POST request to `repos/{owner}/{repo}/actions/workflows/{id}/dispatches`.
3. **Queue**: GitHub queues the job.
4. **Execution**: Runner picks up the job, executes "make build".
5. **Output**: Runner zips the binary and logs, uploading to `actions/artifacts`.
6. **Retrieval**: CLI polls `actions/runs/{id}`, detects "Success", downloads artifact.
7. **Presentation**: CLI unzips and displays logs to user.

## 4. Security
- **Authentication**: GitHub Personal Access Tokens (PAT).
- **Isolation**: Each job runs in an ephemeral VM (Runner) provided by GitHub.
- **Encryption**: Payloads can be encrypted client-side before dispatch (Future Scope).
