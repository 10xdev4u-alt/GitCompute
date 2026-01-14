# ⚡ GitCompute (ActionGrid)

> **The Poor Man's Supercomputer.** 🥘  
> *Turn your GitHub Repository into a High-Performance Distributed Compute Cluster.*

![Status](https://img.shields.io/badge/status-cooking-orange?style=for-the-badge)
![Go](https://img.shields.io/badge/made%20with-Go-blue?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/license-MIT-green?style=for-the-badge)
![Vibe](https://img.shields.io/badge/vibe-100%25-ff69b4?style=for-the-badge)

---

## 🧐 What is this Sorcery?

**Problem:** AWS Lambda is expensive. EC2 is a pain to manage. Your laptop sounds like a jet engine when compiling Rust.

**Solution:** **GitCompute**.
We hijacked GitHub Actions. Instead of just running CI/CD, we treat it as a **Serverless CPU**. You type a command on your potato laptop, and GitHub's massive cloud fleet executes it instantly.

**It's like `ssh` but for GitHub Actions.**

## ✨ Features (The "Mass" Stuff)

- **🔥 Serverless-ish**: No servers. No Kubernetes. Just you, a repo, and pure power.
- **🌍 Cross-Platform**:
  - Want to test a Windows `.exe`? `git-compute run --os windows-latest`
  - Need a Mac? `git-compute run --os macos-latest`
  - Linux? Duh.
- **⚡ Instant Parallelism**: (Coming Soon) Split 1,000 tasks across 20 runners. 20x speedup for free.
- **💸 Free Tier**: Uses your GitHub Actions minutes (2,000 free mins/month for public repos).

## 🚀 Quick Start (Let's Cook!)

### 1. Installation
Get the binary hot out of the oven:
```bash
go install github.com/10xdev4u-alt/cicd-as-a-service@latest
# Note: Rename the binary if needed, or alias it to 'gc' for speed!
```

### 2. Setup (The "One-Time" Thing)
Run the wizard. It asks for your Token, Owner, and Repo.
```bash
git-compute init
```
*Don't have a token? Get one [here](https://github.com/settings/tokens) with `repo` scope.*

### 3. Deploy the Worker
You need the "Receiver" on your repo.
1. Copy the `worker.yml` to your repo's `.github/workflows/` folder.
2. Push it to `main`.
*(Check `docs/` or the source code for the yaml file if you lost it!)*

### 4. Fire Away! 🔫
```bash
# The classic
git-compute run --cmd "echo Hello from the Cloud"

# The "I need a Windows Machine right now"
git-compute run --cmd "systeminfo" --os windows-latest

# The "Watch me build this"
git-compute run --cmd "make build-production" --watch
```

## 🧠 How it Works (The Secret Sauce)

1. **You** type a command.
2. **CLI** sends a `workflow_dispatch` signal to GitHub.
3. **GitHub** wakes up a Runner (VM).
4. **Runner** executes your command and captures stdout/stderr.
5. **Runner** zips the logs and uploads them as an Artifact.
6. **CLI** sees the job finish, downloads the zip, extracts it, and shows you the result.

**Latency?** About 3-10 seconds for VM boot. Not real-time, but fast enough for heavy lifting.

## 📚 Documentation
- [Usage Guide](USAGE.md) - Detailed commands and flags.
- [Architecture (HLD/LLD)](docs/architecture/HLD.md) - For the nerds.
- [Contributing](CONTRIBUTING.md) - Join the kitchen.
- [Future Roadmap](FUTURE_WORKS.md) - What we are cooking next.

## 🤝 Join the Crew
Built by **PrinceTheProgrammer** (10xdev4u-alt).
We are looking for fellow chefs to make this project huge. Check [COLLABORATION.md](COLLABORATION.md).

---
*Disclaimer: Don't mine crypto. GitHub will ban you. Use this for builds, tests, and science.* 🧪