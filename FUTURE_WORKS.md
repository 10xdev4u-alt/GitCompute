# 🔮 Future Works & Roadmap

We have cooked the MVP, but the feast is just beginning. Here is the vision for **GitCompute v2.0**.

## 🚀 Phase 1: The "Data Tunnel" (Enhanced I/O)
- [ ] **File Uploads**: Allow users to upload local files/scripts before execution.
    - *Idea*: Zip local directory -> Upload to temporary GitHub Release/Artifact -> Runner downloads it -> Execute.
- [ ] **Live Streaming**: Stream logs in real-time using GitHub Actions streaming APIs (WebSocket).

## ⚡ Phase 2: The "Grid" (Parallelism)
- [ ] **Matrix Execution**:
    ```bash
    git-compute run --matrix "os=linux,windows,macos" --cmd "make build"
    ```
- [ ] **Sharding**: Split a large input file into chunks and process them in parallel across 10+ runners.
- [ ] **Map-Reduce**: Aggregate results from multiple runners into a single final report.

## 🔒 Phase 3: Enterprise Features
- [ ] **Encryption**: Encrypt payloads client-side so GitHub (and we) cannot see the commands/data.
- [ ] **Private Runners**: Support for self-hosted runners for specialized hardware (GPU).
- [ ] **RBAC**: Role-based access control for who can trigger jobs.

## 🖥️ Phase 4: The Interface
- [ ] **Web Dashboard**: A standalone TUI or Web UI to visualize the "Grid" status.
- [ ] **Prometheus Metrics**: Expose metrics for monitoring compute usage.

---
*Got a crazy idea? Open an Issue or PR!*
