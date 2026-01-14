# GitCompute (ActionGrid)

> **The Poor Man's Supercomputer.**  
> Turn your GitHub Repository into a High-Performance Distributed Compute Cluster.

![Status](https://img.shields.io/badge/status-pre--alpha-red)
![Go](https://img.shields.io/badge/go-1.21+-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## 🔥 The Pitch
AWS Lambda is expensive. EC2 is annoying to manage. **GitCompute** uses GitHub Actions as a serverless CPU and Git as the communication bus.

You have a heavy task (e.g., "Compile this Rust binary for Windows, Linux, and Mac" or "Process these 500 images"). You type one command on your laptop, and GitHub's fleet does the work for you, returning the results instantly.

## 🚀 Features
- **Serverless-ish**: No servers to manage. Just a repo.
- **Cross-Platform**: Compile Windows .exe from a MacBook.
- **Massive Parallelism**: Split 1000 tasks across 20 runners instantly.
- **Free**: Leveraging the GitHub Free Tier (2,000 mins/month) for legitimate build/test tasks.

## 🛠️ Architecture
See [High Level Design](docs/architecture/HLD.md) and [Low Level Design](docs/architecture/LLD.md).

## ⚡ Quick Start

### Prerequisites
1. A GitHub Account.
2. A Personal Access Token (PAT) with `repo` scope.
3. Go installed locally.

### Installation
```bash
go install github.com/10xdev4u-alt/git-compute@latest
```

### Usage
```bash
# Setup the repo
git-compute init

# Run a command on a remote runner
git-compute run --cmd "echo Hello World" --os ubuntu-latest
```

## 👨‍💻 Author
PrinceTheProgrammer (10xdev4u-alt)
