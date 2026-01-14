# 🔧 Troubleshooting Guide

Something not working? Don't panic. Here are the common issues and fixes.

## 1. "401 Unauthorized" or "404 Not Found"

**Symptoms:**
- The CLI crashes immediately when you run a command.
- Error message says `GET https://api.github.com/...: 404`.

**Causes:**
- Your **Personal Access Token (PAT)** is invalid or expired.
- Your PAT does not have the `repo` scope.
- The `owner` or `repo` name in `.git-compute.yaml` is wrong.
- The repository is Private, and your token doesn't have access.

**Fix:**
Run `git-compute init` again and double-check your credentials.

## 2. "Timeout waiting for workflow run to appear"

**Symptoms:**
- CLI says `Triggering command...` (Success).
- Then it hangs on `Waiting for run to start...` for 10-15 seconds.
- Finally errors out.

**Causes:**
- GitHub Actions is having an outage (check status.github.com).
- You dispatched the event, but the `.github/workflows/worker.yml` file **does not exist** on the `main` branch of the remote repo.
- The YAML file has syntax errors (check the Actions tab in your browser).

**Fix:**
- Go to your repo on GitHub.com -> Actions tab.
- Do you see "GitCompute Worker" on the left? If not, you didn't push the YAML correctly.

## 3. "No artifact named 'execution-result' found"

**Symptoms:**
- The job completes successfully (Green).
- The CLI says `Failed to download artifact`.

**Causes:**
- The command you ran failed so catastrophically that it didn't even produce an `output.log`.
- The `worker.yml` is modified and doesn't upload artifacts with the name `execution-result`.

**Fix:**
- Check the logs in the GitHub Browser UI to see why the upload step failed.

## 4. "Illegal file path" (Zip Slip)

**Symptoms:**
- Error during unzipping.

**Causes:**
- The remote job tried to create a file with `../` in the name. This is a security block in our CLI.

**Fix:**
- Ensure your remote scripts output to standard paths.
