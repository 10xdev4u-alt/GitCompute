# Low Level Design (LLD) - GitCompute

## 1. Package Structure

```
/cmd
  /git-compute   # Main entry point
/pkg
  /api           # GitHub API wrappers
  /config        # Viper configuration (auth, defaults)
  /runner        # Logic for dispatching and polling
  /utils         # Zip/Unzip, Logging
```

## 2. Core Structs

### `Config`
```go
type Config struct {
    GitHubToken string `mapstructure:"github_token"`
    Owner       string `mapstructure:"owner"`
    Repo        string `mapstructure:"repo"`
    WorkflowID  string `mapstructure:"workflow_id"`
}
```

### `JobPayload`
The data structure sent to the Action inputs.
```json
{
  "command": "string",
  "os": "ubuntu-latest | windows-latest | macos-latest",
  "timeout": "int"
}
```

### `JobResult`
Metadata retrieved from the completed run.
```go
type JobResult struct {
    RunID      int64
    Status     string
    ArtifactURL string
}
```

## 3. Interfaces

### `Dispatcher`
Abstracts the GitHub API to allow testing/mocking.
```go
type Dispatcher interface {
    TriggerWorkflow(ctx Context, input JobPayload) (int64, error)
    GetRunStatus(ctx Context, runID int64) (string, error)
    DownloadArtifact(ctx Context, runID int64) ([]byte, error)
}
```

## 4. Error Handling
- **API Rate Limits**: Implement exponential backoff.
- **Runner Failures**: Parse workflow logs to distinguish between infrastructure failure vs. script failure.
