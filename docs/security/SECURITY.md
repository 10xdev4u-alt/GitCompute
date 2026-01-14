# 🛡️ Security Model

GitCompute is a powerful tool, but with great power comes great responsibility. Here is how we handle security.

## 1. Authentication
We use **GitHub Personal Access Tokens (PATs)**.
- The CLI reads the token from `~/.git-compute.yaml`.
- This file is stored with user-only read permissions (on Linux/Mac).
- **Recommendation**: Never commit your config file to git.

## 2. Transport Security
All communication happens over **HTTPS** via the official GitHub API (`api.github.com`). We do not run any intermediate servers. It is a direct link between your CLI and GitHub.

## 3. Execution Isolation
Every job runs in a **GitHub Hosted Runner**.
- These are ephemeral Virtual Machines (VMs).
- They are destroyed immediately after the job finishes.
- This ensures clean state and isolation from other jobs.

## 4. Risks & Mitigations

### ⚠️ Risk: Malicious Commands
If someone gets access to your `git-compute` config, they can run any command on your GitHub Actions quota.
- **Mitigation**: Treat your PAT like a password. Rotate it frequently.

### ⚠️ Risk: Secret Leakage
If you run `env` or `printenv` in a command, your GitHub Secrets (if exposed to the workflow) could be printed to the logs and downloaded to your machine.
- **Mitigation**: Be careful what you echo to stdout.

## 5. Future Plans (Encryption)
In v2.0, we plan to implement client-side encryption.
- The CLI will encrypt the command payload using a shared secret.
- The Runner will decrypt it before execution.
- This prevents GitHub (the platform) from easily reading your commands if they inspect the payload.
