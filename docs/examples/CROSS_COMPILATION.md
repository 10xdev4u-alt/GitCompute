# 🧪 Example: Cross-Platform Compilation

One of the best use cases for GitCompute is compiling binaries for operating systems you don't own.

## Scenario
You are on a **MacBook**, but you need to release a `.exe` for **Windows** and a binary for **Linux**.

## The Old Way
1. Install VirtualBox.
2. Install Windows.
3. Install Go on Windows.
4. Build.
5. Repeat for Linux.
6. Cry.

## The GitCompute Way

### 1. Build for Linux
```bash
git-compute run \
  --os ubuntu-latest \
  --cmd "go build -o myapp-linux main.go && ls -lh myapp-linux"
```
*Result: The logs confirm the build. You can modify the worker to upload the binary itself.*

### 2. Build for Windows
```bash
git-compute run \
  --os windows-latest \
  --cmd "go build -o myapp.exe main.go && dir myapp.exe"
```

## Advanced: Retrieve the Binary
Currently, the standard `worker.yml` captures `stdout`. To retrieve the actual **binary**, you would modify your `worker.yml` to zip the binary instead of just the log.

**Modified Step in `worker.yml`:**
```yaml
      - name: Execute and Package
        run: |
          ${{ github.event.inputs.command }}
          # Assume the command created a 'dist' folder
          zip -r result.zip dist/
      
      - uses: actions/upload-artifact@v4
        with:
          name: execution-result
          path: result.zip
```
Now, `git-compute` will download your compiled binaries automatically!
